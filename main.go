package main

import (
	"database/sql"
	"log"

	"github.com/TheDP66/simple_bank_go/api"
	db "github.com/TheDP66/simple_bank_go/db/sqlc"
	"github.com/TheDP66/simple_bank_go/db/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)

	}
}
