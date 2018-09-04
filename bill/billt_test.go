package bill

import (
	"time"

	check "gopkg.in/check.v1"
)

func (s *S) TestCalculateBillOnNormalHours(c *check.C) {
	start := time.Date(2018, 9, 4, 14, 0, 0, 651387237, time.UTC)
	end := time.Date(2018, 9, 4, 20, 0, 0, 651387237, time.UTC)
	toPay, err := Calculate(start, end)
	c.Assert(err, check.Equals, nil)
	c.Assert(toPay, check.Equals, 32.76)
}

func (s *S) TestCalculateBillOnInvalidEndTime(c *check.C) {
	start := time.Date(2018, 9, 4, 20, 0, 0, 651387237, time.UTC)
	end := time.Date(2018, 9, 4, 14, 0, 0, 651387237, time.UTC)
	_, err := Calculate(start, end)
	c.Assert(err, check.NotNil)
}
