package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)
func GetContext() context.Context{
	return context.Background()
}
func TestMain(m* testing.M){
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}

// func TestMain(m *testing.M){
// 	conn, err := sql.Open(dbDriver,dbSource)
// 	if err != nil {
// 		log.Fatal("cannot connect to db: ", err)
// 	}

// 	testQueries = New(conn)
// 	os.Exit(m.Run())
// }