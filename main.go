package main

import (
	"context"
	"log"

	"github.com/Kalpesh-07/GoBank/api"
	db "github.com/Kalpesh-07/GoBank/db/sqlc"
	"github.com/Kalpesh-07/GoBank/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config")
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}

	store := db.NewStore(connPool)
	runGinServer(config, store)
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server ", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot create server ", err)
	}
}
