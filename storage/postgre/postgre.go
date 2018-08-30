package postgre

import (
	"database/sql"
	"fmt"
	"os"

	// pq is the postgres driver
	_ "github.com/lib/pq"
)

type Postgre struct {
	db *sql.DB
}

func New() (postg *Postgre, err error) {
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

func SaveCall() error {
	return nil
}
