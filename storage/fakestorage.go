package storage

import (
	"github.com/arxdsilva/olist/record"
)

type FakeStorage struct {
	records []record.Record
}

func (f FakeStorage) SaveRecord(r record.Record) (err error) {
	f.records = append(f.records, r)
	return
}
