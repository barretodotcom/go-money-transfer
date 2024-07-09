package database

import (
	"context"

	"github.com/go-money-transfer/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

var conn *pgxpool.Pool
var err error

func Connect(config config.Config) error {
	conn, err = pgxpool.Connect(context.Background(), config.DatabaseURL)

	return err
}

func GetConnection() *pgxpool.Pool {
	return conn
}
