package storage

import (
	"github.com/arxdsilva/olist/bill"
	"github.com/arxdsilva/olist/record"
)

type FakeStorage struct {
	records []record.Record
}

func (f FakeStorage) SaveRecord(r record.Record) (err error) {
	f.records = append(f.records, r)
	return
}

func (f FakeStorage) UUIDFromStart(r record.Record) (uuid string, err error) {
	return
}

func (f FakeStorage) BillFromID(id string) (b bill.Bill, err error) {
	return
}

func (f FakeStorage) CallsFromBillID(id string) (cs []bill.Call, err error) {
	return
}

func (f FakeStorage) RecordsFromBill(b bill.Bill) (rs []record.Record, err error) {
	return
}

func (f FakeStorage) SaveBill(b bill.Bill) (err error) {
	return
}

func (f FakeStorage) SaveCalls(c []bill.Call) (err error) {
	return
}
