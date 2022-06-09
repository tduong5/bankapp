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
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource) // returns connection obj and error
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB) // else create new testQueries obj

	os.Exit(m.Run())
}
