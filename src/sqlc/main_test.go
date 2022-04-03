package db

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var db *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://admin:P@ssw0rd@localhost:5432/simple_bank?sslmode=disable"
)

var ctx = context.Background()

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to database!")
	}

	db = New(conn)

	os.Exit(m.Run())
}
