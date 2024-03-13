package main

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	data "github.com/maknop/pokedex-web-app/data"
	db "github.com/maknop/pokedex-web-app/database"
	pokemon "github.com/maknop/pokedex-web-app/routes"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
	router.Static("/styles", "./styles")

	time.Sleep(10 * time.Second)
	go db.EstablishConnection()

	allPokemon := data.GetPokemonData()

	log.Info("pokedex API running...")

	// Routes
	pokemon.LandingPage(router, allPokemon)
	pokemon.PokemonById(router, allPokemon)

	router.Run("localhost:8080")
}
