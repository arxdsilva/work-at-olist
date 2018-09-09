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
