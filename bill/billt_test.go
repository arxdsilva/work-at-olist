package bill

import (
	"time"

	check "gopkg.in/check.v1"
)

func (s *S) TestCalculateBillOnNormalHours(c *check.C) {
	start := time.Date(2018, 9, 4, 14, 0, 0, 0, time.UTC)
	end := time.Date(2018, 9, 4, 20, 0, 0, 0, time.UTC)
	toPay, err := Calculate(start, end)
	c.Assert(err, check.Equals, nil)
	c.Assert(toPay, check.Equals, 32.76)
}

func (s *S) TestCalculateBillOnInvalidEndTime(c *check.C) {
	start := time.Date(2018, 9, 4, 20, 0, 0, 0, time.UTC)
	end := time.Date(2018, 9, 4, 14, 0, 0, 0, time.UTC)
	_, err := Calculate(start, end)
	c.Assert(err, check.NotNil)
}

func (s *S) TestCalculateBillOnLowerRateHours(c *check.C) {
	start := time.Date(2018, 9, 4, 22, 1, 0, 0, time.UTC)
	end := time.Date(2018, 9, 5, 5, 1, 0, 0, time.UTC)
	toPay, err := Calculate(start, end)
	c.Assert(err, check.IsNil)
	c.Assert(toPay, check.Equals, 0.36)
}

func (s *S) TestCalculateBillOnMixedHours(c *check.C) {
	start := time.Date(2018, time.September, 4, 21, 0, 0, 0, time.UTC)
	end := time.Date(2018, time.September, 4, 22, 30, 0, 0, time.UTC)
	toPay, err := Calculate(start, end)
	c.Assert(err, check.IsNil)
	c.Assert(toPay, check.Equals, 5.76)
}
