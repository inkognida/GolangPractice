package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func main() {
	databaseUrl := "postgres://hardella:123@localhost:5432/postgres"

	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalln("Unable to connect to DB")
	}
	defer dbPool.Close()

}
