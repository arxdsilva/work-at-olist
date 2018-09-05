package bill

import (
	"errors"
	"time"
)

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
	if (startHour >= 22) && (endHour <= 6) {
		return 0.36, err
	}
	var unbillableTime float64
	switch {
	case startsOnUnbillableHours(start):
		unbillableTime = unbillableTimeAtStart(start, end)
	}
	timeUsed := end.Sub(start)
	minutesUsed := timeUsed.Minutes()
	r = 0.36 + (0.09 * (minutesUsed - unbillableTime))
	return r, err
}

func startsOnUnbillableHours(start time.Time) bool {
	unbillableTimeStart := time.Date(start.Year(), start.Month(), start.Day(), 22, 0, 0, 0, time.UTC)
	unbillableTimeEnd := time.Date(start.Year(), start.Month(), start.Day(), 6, 0, 0, 0, time.UTC)
	if start.After(unbillableTimeStart) {
		return true
	} else if start.Before(unbillableTimeEnd) {
		return true
	}
	return false
}
