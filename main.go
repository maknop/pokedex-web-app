package main

import (
	"fmt"

	"github.com/maknop/pokedex-app/api/handler"
	log "github.com/sirupsen/logrus"
)

func main() {
	data, err := handler.GetPokemon()
	if err != nil {
		log.Error("there was an issue retrieving data.")
	}

	fmt.Println(string(data))
}
