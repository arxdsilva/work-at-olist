package storage

import (
	"github.com/arxdsilva/olist/bill"
	"github.com/arxdsilva/olist/record"
)

type Storage interface {
	SaveRecord(record.Record) error
	UUIDFromStart(record.Record) (string, error)
	BillFromID(string) (bill.Bill, error)
	CallsFromBillID(string) ([]bill.Call, error)
	RecordsFromBill(bill.Bill) ([]record.Record, error)
}
