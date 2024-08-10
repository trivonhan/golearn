package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:postgres@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	x

	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
