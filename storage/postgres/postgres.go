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
	return
}

func (p Postgres) UUIDFromStart(r record.Record) (uuid string, err error) {
	query := "select id from records where call_id = $1"
	row := p.db.QueryRow(query, r.CallID)
	err = row.Scan(&uuid)
	return uuid, err
}

func (p Postgres) BillFromID(id string) (b bill.Bill, err error) {
	query := "select bill_id, sub_number, b_month, b_year from bills where bill_id = $1"
	return b, p.db.QueryRow(query, id).Scan(&b.ID, &b.SubscriberNumber, &b.Month, &b.Year)
}

func (p Postgres) CallsFromBillID(id string) (cs []bill.Call, err error) {
	query := "select destination, start_date, start_time, duration, price from calls where bill_id = $1"
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
	query := "select r_type, time_stamp, r_source, destination, call_id, r_month, id from records where r_month = $1 and r_year = $2 and r_source = $3"
	rows, err := p.db.Query(query, b.Month, b.Year, b.SubscriberNumber)
	if err != nil {
		return
	}
	var buff []record.Record
	defer rows.Close()
	for rows.Next() {
		r := new(record.Record)
		if err = rows.Scan(&r.Type, &r.TimeStamp, &r.Source, &r.Destination, &r.CallID, &r.Month, &r.ID); err != nil {
			return
		}
		rs = append(rs, *r)
		buff = append(buff, *r)
	}
	for _, r := range buff {
		var endr record.Record
		query = "select r_type, time_stamp, r_source, destination, call_id, r_month from records where r_type = 'end' and r_month = $1 and r_year = $2 and id = $3"
		row := p.db.QueryRow(query, b.Month, b.Year, r.ID)
		err = row.Scan(&endr.Type, &endr.TimeStamp, &endr.Source, &endr.Destination, &endr.CallID, &endr.Month)
		if err != nil {
			return
		}
		rs = append(rs, endr)
	}
	return
}

func (p Postgres) SaveBill(b bill.Bill) (err error) {
	query := "insert into bills (bill_id, b_month, b_year, sub_number, total) values ($1, $2, $3, $4, $5)"
	_, err = p.db.Exec(query, b.ID, b.Month, b.Year, b.SubscriberNumber, b.Total)
	return
}

func (p Postgres) SaveCalls(calls []bill.Call) (err error) {
	for _, call := range calls {
		query := "insert into calls (bill_id, duration, price, start_date, start_time, destination) values ($1, $2, $3, $4, $5, $6)"
		_, err = p.db.Exec(query, call.BillID, call.CallDuration, call.CallPrice, call.CallStartDate, call.CallStartTime, call.Destination)
		if err != nil {
			return
		}
	}
	return
}
