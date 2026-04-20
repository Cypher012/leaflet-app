package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB(url string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatal(err)
	}
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected..")

	return pool
}
