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

func (s *S) TestStartsAfterUnbillableHours(c *check.C) {
	start := time.Date(2018, time.September, 4, 23, 0, 0, 0, time.UTC)
	c.Assert(startsOnUnbillableHours(start), check.Equals, true)
}

func (s *S) TestStartsAfterUnbillableHoursNextDay(c *check.C) {
	start := time.Date(2018, time.September, 4, 2, 0, 0, 0, time.UTC)
	c.Assert(startsOnUnbillableHours(start), check.Equals, true)
}

func (s *S) TestStartsAfterUnbillableHoursMustBeFalse(c *check.C) {
	start := time.Date(2018, time.September, 4, 21, 0, 0, 0, time.UTC)
	c.Assert(startsOnUnbillableHours(start), check.Equals, false)
}

func (s *S) TestUnbillableTimeAtStart(c *check.C) {
	start := time.Date(2018, time.September, 4, 22, 0, 0, 0, time.UTC)
	end := time.Date(2018, time.September, 4, 22, 30, 0, 0, time.UTC)
	unbTime := unbillableTimeAtStart(start, end)
	c.Assert(unbTime, check.Equals, float64(30))
}

func (s *S) TestUnbillableTimeAtStartWithEndOnNextDay(c *check.C) {
	start := time.Date(2018, time.September, 4, 23, 0, 0, 0, time.UTC)
	end := time.Date(2018, time.September, 5, 0, 30, 0, 0, time.UTC)
	unbTime := unbillableTimeAtStart(start, end)
	c.Assert(unbTime, check.Equals, float64(90))
}

func (s *S) TestUnbillableTimeAtStartWithEndAfterUnbillableHours(c *check.C) {
	start := time.Date(2018, time.September, 4, 5, 0, 0, 0, time.UTC)
	end := time.Date(2018, time.September, 4, 6, 30, 0, 0, time.UTC)
	unbTime := unbillableTimeAtStart(start, end)
	c.Assert(unbTime, check.Equals, float64(60))
}

func (s *S) TestEndsOnUnbillableHours(c *check.C) {
	end := time.Date(2018, time.September, 4, 23, 0, 0, 0, time.UTC)
	c.Assert(endsOnUnbillableHours(end), check.Equals, true)
}

func (s *S) TestEndsOnUnbillableHoursNextDay(c *check.C) {
	end := time.Date(2018, time.September, 4, 2, 0, 0, 0, time.UTC)
	c.Assert(endsOnUnbillableHours(end), check.Equals, true)
}

func (s *S) TestEndsOnUnbillableHoursMustBeFalse(c *check.C) {
	end := time.Date(2018, time.September, 4, 21, 0, 0, 0, time.UTC)
	c.Assert(endsOnUnbillableHours(end), check.Equals, false)
}

func (s *S) TestUnbillableTimeAtEnd(c *check.C) {
	start := time.Date(2018, time.September, 4, 21, 0, 0, 0, time.UTC)
	end := time.Date(2018, time.September, 4, 22, 30, 0, 0, time.UTC)
	unbTime := unbillableTimeAtEnd(start, end)
	c.Assert(unbTime, check.Equals, float64(30))
}

func (s *S) TestUnbillableTimeAtEndWithEndOnNextDay(c *check.C) {
	start := time.Date(2018, time.September, 4, 21, 0, 0, 0, time.UTC)
	end := time.Date(2018, time.September, 5, 0, 30, 0, 0, time.UTC)
	unbTime := unbillableTimeAtEnd(start, end)
	c.Assert(unbTime, check.Equals, float64(150))
}

func (s *S) TestUnbHoursBetweenStartAndEnd(c *check.C) {
	start := time.Date(2018, time.September, 4, 21, 0, 0, 0, time.UTC)
	end := time.Date(2018, time.September, 5, 7, 0, 0, 0, time.UTC)
	c.Assert(unbHoursBetweenStartAndEnd(start, end), check.Equals, true)
}
