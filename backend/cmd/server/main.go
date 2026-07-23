package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		connString = "postgres://linuquest_user:password@localhost:5432/linuquest"
	}

	db, err := pgxpool.New(context.Background(), connString)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.Run(":8080")

}
