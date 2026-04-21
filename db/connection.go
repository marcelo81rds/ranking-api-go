package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func Connect() error {
	var err error

	databaseUrl := os.Getenv("DATABASE_URL")

	Conn, err = pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		return err
	}

	fmt.Println("Conectado ao PostgreSQL!")
	return nil
}
