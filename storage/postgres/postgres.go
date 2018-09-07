package postgres

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/arxdsilva/olist/record"

	// pq is the postgres driver
	_ "github.com/lib/pq"
)

type Postgre struct {
	db *sql.DB
}

func New() (postg Postgre, err error) {
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbconfig := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", user, dbname, password, host)
	db, err := sql.Open("postgres", dbconfig)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	postg.db = db
	return
}

func (p Postgre) SaveRecord(r record.Record) (err error) {
	query := "insert into records (id, r_type, time_stamp, call_id, r_source, destination) values ($1, $2, $3, $4, $5, $6)"
	_, err = p.db.Exec(query, r.ID, r.Type, r.TimeStamp, r.CallID, r.Source, r.Destination)
	return err
}
