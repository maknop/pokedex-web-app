package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect() error {
	connStr := "user=ADMIN dbname=pokemon sslmode=disable"
	_, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("connection failed: %s\n", err)
	}

	fmt.Print("connection to postgres database was successful!")

	return nil
}
