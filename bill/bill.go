package bill

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

type Bill struct {
	ID               string  `json:"id"`
	SubscriberNumber string  `json:"subscriber"`
	Month            string  `json:"month"` // only after it has ended (not current month)
	Year             string  `json:"year"`
	Calls            []Call  `json:"calls"`
	Total            float64 `json:"total"`
}

func New(month, year, subNumber string) (b Bill) {
	b.ID = fmt.Sprintf("%s%s%s", subNumber, month, year)
	b.Month = month
	b.Year = year
	b.SubscriberNumber = subNumber
	return
}

func (b *Bill) CalculateTotalPrice() {
	var total float64
	for _, v := range b.Calls {
		total += v.CallPrice
	}
	b.Total = total
	return
}

// Calculate makes the math
// of the time to be charged by the telephone records
// 600 > 2200: 0.36 + 0.09/min
// 2200 > 600: 0.36
func Calculate(start, end time.Time) (r float64, err error) {
	if inv := start.Sub(end); inv.Minutes() > 0 {
		return r, errors.New("Start time cannot be after the call has ended")
	}
	startHour := start.Hour()
	endHour := end.Hour()
	if (startHour >= 22) && (endHour <= 6) && (start.Day() != end.Day()) {
		return 0.36, err
	}
	var unbillableTime float64
	switch {
	case startsOnUnbillableHours(start):
		unbillableTime = unbillableTimeAtStart(start, end)
	case endsOnUnbillableHours(end):
		unbillableTime = unbillableTimeAtEnd(start, end)
	case unbHoursBetweenStartAndEnd(start, end):
		unbillableTime = unbillableTimeBetweenSE(start, end)
	}
	timeUsed := end.Sub(start)
	minutesUsed := timeUsed.Minutes()
	minRate, err := minuteRate()
	if err != nil {
		return
	}
	baseRate, err := baseRate()
	if err != nil {
		return
	}
	fullMinutesUsed := float64(int(minutesUsed - unbillableTime))
	r = round(baseRate+(minRate*(fullMinutesUsed)), 100)
	return r, err
}

func startsOnUnbillableHours(start time.Time) bool {
	unbillableTimeStart := time.Date(start.Year(), start.Month(), start.Day(), 21, 59, 59, 0, time.UTC)
	unbillableTimeEnd := time.Date(start.Year(), start.Month(), start.Day(), 6, 1, 0, 0, time.UTC)
	if start.After(unbillableTimeStart) || start.Before(unbillableTimeEnd) {
		return true
	}
	return false
}

func unbillableTimeAtStart(s, e time.Time) (unbillableTime float64) {
	var unbillableTimeEnd time.Time
	unbillableTimeStart := time.Date(s.Year(), s.Month(), s.Day(), 21, 59, 59, 0, time.UTC)
	if s.After(unbillableTimeStart) {
		unbillableTimeEnd = time.Date(s.Year(), s.Month(), s.Day()+1, 6, 0, 0, 0, time.UTC)
	} else {
		unbillableTimeEnd = time.Date(s.Year(), s.Month(), s.Day(), 6, 0, 0, 0, time.UTC)
	}
	if e.Before(unbillableTimeEnd) {
		unbillableDuration := e.Sub(s)
		return unbillableDuration.Minutes()
	}
	unbillableDuration := unbillableTimeEnd.Sub(s)
	return unbillableDuration.Minutes()
}

func endsOnUnbillableHours(end time.Time) bool {
	unbillableTimeStart := time.Date(end.Year(), end.Month(), end.Day(), 21, 59, 59, 0, time.UTC)
	unbillableTimeEnd := time.Date(end.Year(), end.Month(), end.Day(), 6, 1, 0, 0, time.UTC)
	if end.After(unbillableTimeStart) || end.Before(unbillableTimeEnd) {
		return true
	}
	return false
}

func unbillableTimeAtEnd(s, e time.Time) (unbillableTime float64) {
	unbillableStartTime := time.Date(e.Year(), e.Month(), s.Day(), 22, 0, 0, 0, time.UTC)
	var unbillableDate time.Time
	if s.Day() != e.Day() {
		unbillableDate = time.Date(e.Year(), e.Month(), e.Day(), 6, 0, 0, 0, time.UTC)
	} else {
		unbillableDate = unbillableStartTime
	}
	if e.Before(unbillableDate) {
		unbillableDuration := e.Sub(unbillableStartTime)
		return unbillableDuration.Minutes()
	}
	unbillableDuration := e.Sub(unbillableDate)
	return unbillableDuration.Minutes()
}

func unbHoursBetweenStartAndEnd(s, e time.Time) bool {
	unbillableTimeStart := time.Date(s.Year(), s.Month(), s.Day(), 22, 0, 0, 0, time.UTC)
	unbillableTimeEnd := time.Date(e.Year(), e.Month(), e.Day(), 6, 0, 0, 0, time.UTC)
	if e.After(unbillableTimeEnd) && s.Before(unbillableTimeStart) && (s.Day() != e.Day()) {
		return true
	}
	return false
}

func unbillableTimeBetweenSE(s, e time.Time) (u float64) {
	days := e.Day() - s.Day()
	u += float64(days * 8 * 60)
	unbillableTimeEnd := time.Date(e.Year(), e.Month(), e.Day(), 22, 0, 0, 0, time.UTC)
	if e.After(unbillableTimeEnd) {
		duration := e.Sub(unbillableTimeEnd)
		u += duration.Minutes()
	}
	return
}

func round(x, unit float64) float64 {
	return math.Round(x*unit) / unit
}

func minuteRate() (rate float64, err error) {
	r := os.Getenv("MIN_RATE")
	if r == "" {
		return 0.09, err
	}
	return strconv.ParseFloat(r, 32)
}

func baseRate() (rate float64, err error) {
	r := os.Getenv("BASE_RATE")
	if r == "" {
		return 0.36, err
	}
	return strconv.ParseFloat(r, 32)
}
