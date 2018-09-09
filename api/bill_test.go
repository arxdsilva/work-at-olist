package api

import (
	"github.com/arxdsilva/olist/record"
	check "gopkg.in/check.v1"
)

func (s *S) Test_filterRecordsPeriod(c *check.C) {
	r := []record.Record{
		record.Record{Type: "start", CallID: "123", Month: 1},
		record.Record{Type: "end", CallID: "123", Month: 1},
	}
	gotRFiltered := filterRecordsPeriod(r, 1)
	c.Assert(len(gotRFiltered), check.Equals, 1)
}

func (s *S) Test_filterRecordsPeriodFilteringRecord(c *check.C) {
	r := []record.Record{
		record.Record{Type: "start", CallID: "123", Month: 1},
		record.Record{Type: "end", CallID: "123", Month: 1},
		record.Record{Type: "start", CallID: "124", Month: 1},
		record.Record{Type: "end", CallID: "124", Month: 2},
	}
	gotRFiltered := filterRecordsPeriod(r, 1)
	c.Assert(len(gotRFiltered), check.Equals, 1)
}

func (s *S) Test_callFromRecords(c *check.C) {
	records := []record.Record{
		record.Record{
			Type:        "start",
			TimeStamp:   "2016-02-29T14:00:00Z",
			Destination: "2199999999",
		},
		record.Record{
			Type:        "end",
			TimeStamp:   "2016-02-29T14:03:00Z",
			Destination: "2199999999",
		},
	}
	call, err := callFromRecords(records)
	c.Assert(err, check.IsNil)
	c.Assert(call.CallPrice, check.Equals, 0.63)
	c.Assert(call.CallDuration, check.Equals, "3m0s")
	c.Assert(call.CallStartDate, check.Equals, 29)
	c.Assert(call.CallStartTime, check.Equals, "2:00PM")
	c.Assert(call.Destination, check.Equals, "2199999999")
}

func (s *S) Test_callsFromRecordsForWrongPeriod(c *check.C) {
	records := []record.Record{
		record.Record{
			CallID:      "123",
			Type:        "start",
			TimeStamp:   "2016-02-29T14:00:00Z",
			Destination: "2199999999",
			Month:       2,
		},
		record.Record{
			CallID:      "123",
			Type:        "end",
			TimeStamp:   "2016-02-29T14:03:00Z",
			Destination: "2199999999",
			Month:       2,
		},
	}
	calls, err := callsFromRecords(records, 1, "")
	c.Assert(err, check.IsNil)
	c.Assert(len(calls), check.Equals, 0)
}

func (s *S) Test_callsFromRecordsSingleRecord(c *check.C) {
	records := []record.Record{
		record.Record{
			CallID:      "123",
			Type:        "start",
			TimeStamp:   "2016-02-29T14:00:00Z",
			Destination: "2199999999",
			Month:       2,
		},
		record.Record{
			CallID:      "123",
			Type:        "end",
			TimeStamp:   "2016-02-29T14:03:00Z",
			Destination: "2199999999",
			Month:       2,
		},
	}
	calls, err := callsFromRecords(records, 2, "")
	c.Assert(err, check.IsNil)
	c.Assert(len(calls), check.Equals, 1)
}

func (s *S) Test_callsFromRecordsMultipleRecords(c *check.C) {
	records := []record.Record{
		record.Record{
			CallID:      "123",
			Type:        "start",
			TimeStamp:   "2016-02-29T14:00:00Z",
			Destination: "2199999999",
			Month:       2,
		},
		record.Record{
			CallID:      "123",
			Type:        "end",
			TimeStamp:   "2016-02-29T14:03:00Z",
			Destination: "2199999999",
			Month:       2,
		},
		record.Record{
			CallID:      "124",
			Type:        "start",
			TimeStamp:   "2016-02-03T14:00:00Z",
			Destination: "2199999999",
			Month:       2,
		},
		record.Record{
			CallID:      "124",
			Type:        "end",
			TimeStamp:   "2016-02-03T14:03:00Z",
			Destination: "2199999999",
			Month:       2,
		},
	}
	calls, err := callsFromRecords(records, 2, "")
	c.Assert(err, check.IsNil)
	c.Assert(len(calls), check.Equals, 2)
}
