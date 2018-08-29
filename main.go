package main

import (
	"github.com/arxdsilva/olist/api"
	"github.com/arxdsilva/olist/storage/postgre"
	"log"
)

func main() {
	s := api.New()
	s.Storage, err := postgre.New()
	if err != nil {
		log.Fatal(err)
	}
	s.Listen()
}
