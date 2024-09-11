package main

import (
	db "github.com/maknop/pokedex-app/internal/database"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Establish connection to postgres database
	err := db.Connect()
	if err != nil {
		log.Error("Failed to connect to postgres database: %w", err)
	}

	// data, err := handler.GetPokemon()
	// if err != nil {
	// 	log.Error("there was an issue retrieving data.")
	// }

	// fmt.Println(string(data))
}
