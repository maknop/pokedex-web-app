package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	log "github.com/sirupsen/logrus"
)

var db *sql.DB

func EstablishConnection() {
	log.Info("setting database environment variables")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDatabase := os.Getenv("POSTGRES_DB")
	postgresHost := os.Getenv("POSTGRES_HOST")

	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s", postgresUser, postgresDatabase, postgresPassword, postgresHost))
	if err != nil {
		log.Fatalf(fmt.Sprintf("connection to postgres container failed: %s", err))
	}

	defer db.Close()

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully Connected")
}
