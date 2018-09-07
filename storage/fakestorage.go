package storage

import (
	"github.com/arxdsilva/olist/record"
)

type FakeStorage struct {
	records []record.Record
}

func (f *FakeStorage) SaveRecords(r record.Record) (err error) {
	f.records = append(f.records, r)
	return
}
