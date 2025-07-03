package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	db "github.com/dargoz/day04/data/local/internal/db"
	_ "github.com/lib/pq" // PostgreSQL driver
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres@localhost:5432/postgres?sslmode=disable"
)

var testQueries *db.Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = db.New(conn)
	os.Exit(m.Run())
}
