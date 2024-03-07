package database

import (
	"fmt"
	"log"
	"os"

	models "github.com/maknop/pokedex-web-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDatabase := os.Getenv("POSTGRES_DB")
	postgresHost := os.Getenv("POSTGRES_HOST")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", postgresHost, postgresUser, postgresPassword, postgresDatabase)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Pokemon{})

	return db
}
