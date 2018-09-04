package bill

import (
	"time"
)

// Calculate makes the math
// of the time to be charged by the telephone records
// 600 > 2200: 0.36 + 0.09/min
// 2200 > 600: 0.36
func Calculate(start, end time.Time) (r float64, err error) {
	timeUsed := end.Sub(start)
	minutesUsed := timeUsed.Minutes()
	r = 0.36 + (0.09 * minutesUsed)
	return r, err
}
