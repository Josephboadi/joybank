package main

import (
	"database/sql"
	"log"

	"github.com/josephboadi/joybank/api"
	db "github.com/josephboadi/joybank/db/sqlc"
	"github.com/josephboadi/joybank/util"
	_ "github.com/lib/pq"
)

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://root:secret@localhost:5434/joy_bank?sslmode=disable"
// 	serverAddress = "0.0.0.0:8080"
// )


func main(){
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBStore)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}

}