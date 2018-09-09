package postgres

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/arxdsilva/olist/bill"
	"github.com/arxdsilva/olist/record"

	// pq is the postgres driver
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func New() (postg Postgres, err error) {
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

func (p Postgres) SaveRecord(r record.Record) (err error) {
	query := "insert into records (id, r_type, time_stamp, call_id, r_source, destination, r_month, r_year) values ($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err = p.db.Exec(query, r.ID, r.Type, r.TimeStamp, r.CallID, r.Source, r.Destination, r.Month, r.Year)
	return err
}

func (p Postgres) UUIDFromStart(r record.Record) (uuid string, err error) {
	query := "select id from records where call_id = $1"
	row := p.db.QueryRow(query, r.CallID)
	err = row.Scan(&uuid)
	return uuid, err
}

func (p Postgres) BillFromID(id string) (b bill.Bill, err error) {
	query := "select id, sub_number, b_month, b_year from bills where bill_id = $1"
	return b, p.db.QueryRow(query, id).Scan(&b.ID, &b.SubscriberNumber, &b.Month, &b.Year)
}

func (p Postgres) CallsFromBillID(id string) (cs []bill.Call, err error) {
	query := "select destination, start_date, start_time, c_duration, c_price from calls where bill_id = $1"
	rows, err := p.db.Query(query, id)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		c := new(bill.Call)
		if err = rows.Scan(&c.Destination, &c.CallStartDate, &c.CallStartTime, &c.CallDuration, &c.CallPrice); err != nil {
			return
		}
		cs = append(cs, *c)
	}
	return
}

func (p Postgres) RecordsFromBill(b bill.Bill) (rs []record.Record, err error) {
	query := "select r_type, time_stamp, r_source, destination from records where r_month = $1 and r_year = $2 and r_source = $3"
	rows, err := p.db.Query(query, b.Month, b.Year, b.SubscriberNumber)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		r := new(record.Record)
		if err = rows.Scan(&r.Type, &r.TimeStamp, &r.Source, &r.Destination); err != nil {
			return
		}
		rs = append(rs, *r)
	}
	return
}
