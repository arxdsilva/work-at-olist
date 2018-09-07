package storage

import (
	"github.com/arxdsilva/olist/record"
)

type Storage interface {
	SaveRecord(record.Record) error
}
