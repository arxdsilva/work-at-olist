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
	query := "insert into records (id, r_type, time_stamp, call_id, r_source, destination, month) values ($1, $2, $3, $4, $5, $6, $7)"
	_, err = p.db.Exec(query, r.ID, r.Type, r.TimeStamp, r.CallID, r.Source, r.Destination, r.Month)
	return err
}

func (p Postgre) UUIDFromStart(r record.Record) (uuid string, err error) {
	query := "select id from records where r_source = $1 and destination = $2 and month = $3 and year = $4 and r_type = $5"
	rows, err := p.db.Query(query, r.Source, r.Destination, r.Month, r.Year, "start")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		if rows.Next() {
			continue
		}
		if err = rows.Scan(&uuid); err != nil {
			return
		}
	}
	return
}
