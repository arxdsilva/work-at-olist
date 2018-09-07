package main

import (
	"log"

	"github.com/arxdsilva/olist/api"
	"github.com/arxdsilva/olist/storage/postgres"
)

func main() {
	s := api.New()
	storage, err := postgres.New()
	if err != nil {
		log.Fatal(err)
	}
	s.Storage = storage
	s.Listen()
}
