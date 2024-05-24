package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/josephboadi/joybank/util"
)

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://root:secret@localhost:5434/joy_bank?sslmode=disable"
// )

var testStore Store
var testDB *sql.DB


func TestMain(m * testing.M){
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	testStore = NewStore(connPool)
	
	os.Exit(m.Run())
}