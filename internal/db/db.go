package datastore

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
)

func main() {
	db, err := newDB()
	if err != nil {
		fmt.Println("newDB", err)
	}
	defer func() {
		_ = db.Close()
		fmt.Println("closed")
	}()

	db.BeginTx(context.Background(), nil)
}

func newDB() (*sql.DB, error) {
	dsn := url.URL{
		Scheme: "postgres",
		Host:   "localhost:5432",
		User:   url.UserPassword("user", "password"),
		Path:   "dbname",
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")

	dsn.RawQuery = q.Encode()

	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		return nil, fmt.Errorf("sql.Open w", err)
	}
	return db, nil

}
