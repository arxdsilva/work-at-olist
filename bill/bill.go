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
	if (end.Hour() >= 22) || (end.Hour() <= 6) {
		u := time.Date(end.Year(), end.Month(), end.Day(), 22, 0, 0, 0, time.UTC)
		endDuration := end.Sub(u)
		unbillableTime = endDuration.Minutes()
	}
	timeUsed := end.Sub(start)
	minutesUsed := timeUsed.Minutes()
	r = 0.36 + (0.09 * (minutesUsed - unbillableTime))
	return r, err
}
