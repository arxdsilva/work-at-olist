package storage

import (
	"errors"

	"github.com/arxdsilva/olist/bill"
	"github.com/arxdsilva/olist/record"
)

type FakeStorage struct {
	records []record.Record
	bills   []bill.Bill
	calls   []bill.Call
}

func (f FakeStorage) SaveRecord(r record.Record) (err error) {
	f.records = append(f.records, r)
	return
}

func (f FakeStorage) UUIDFromStart(r record.Record) (uuid string, err error) {
	return
}

func (f FakeStorage) BillFromID(id string) (b bill.Bill, err error) {
	return b, errors.New("Not found")
}

func (f FakeStorage) CallsFromBillID(id string) (cs []bill.Call, err error) {
	return
}

func (f FakeStorage) RecordsFromBill(b bill.Bill) (rs []record.Record, err error) {
	return
}

func (f FakeStorage) SaveBill(b bill.Bill) (err error) {
	f.bills = append(f.bills, b)
	return
}

func (f FakeStorage) SaveCalls(c []bill.Call) (err error) {
	for _, call := range c {
		f.calls = append(f.calls, call)
	}
	return
}
