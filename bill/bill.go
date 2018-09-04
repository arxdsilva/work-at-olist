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
	timeUsed := end.Sub(start)
	minutesUsed := timeUsed.Minutes()
	r = 0.36 + (0.09 * minutesUsed)
	return r, err
}
