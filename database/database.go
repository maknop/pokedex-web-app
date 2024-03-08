package database

import (
	"fmt"
	"os"

	models "github.com/maknop/pokedex-web-app/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect - Establish connection to postgres database
func Connect() *gorm.DB {
	log.Info("setting database environment variables")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDatabase := os.Getenv("POSTGRES_DB")
	postgresHost := os.Getenv("POSTGRES_HOST")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", postgresHost, postgresUser, postgresPassword, postgresDatabase)

	log.Info("setting up connection to postgresql database")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Pokemon{})

	return db
}
