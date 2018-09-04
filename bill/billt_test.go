package bill

import (
	"time"

	check "gopkg.in/check.v1"
)

func (s *S) TestCalculateBillOnNormalHours(c *check.C) {
	start := time.Date(2018, 9, 4, 14, 0, 0, 651387237, time.UTC)
	end := time.Date(2018, 9, 4, 20, 0, 0, 651387237, time.UTC)
	toPay := Calculate(start, end)
	c.Assert(toPay, check.Equals, 2)
}
