package main

import (
	"github.com/maknop/pokedex-app/routes"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("vim-go")
	routes.GetPokemon()
}
