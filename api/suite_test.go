package api

import (
	"testing"

	"github.com/arxdsilva/olist/storage"
	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type S struct {
	db *storage.Storage
}

var _ = check.Suite(&S{})
