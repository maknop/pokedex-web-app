package main

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	data "pokedex-web-app/data"
	pokemon "pokedex-web-app/routes"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
	router.Static("/styles", "./styles")

	allPokemon := data.GetPokemonData()

	log.Info("Pokedex API running...")

	// Routes
	pokemon.LandingPage(router, allPokemon)
	pokemon.PokemonById(router, allPokemon)

	router.Run("localhost:8080")

	time.Sleep(5 * time.Second)
}
