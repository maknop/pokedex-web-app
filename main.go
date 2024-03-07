package main

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	data "github.com/maknop/pokedex-web-app/data"
	pokemon "github.com/maknop/pokedex-web-app/routes"
	db "github.com/maknop/pokedex-web-app/database"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
	router.Static("/styles", "./styles")

	db.Connect()

	allPokemon := data.GetPokemonData()

	log.Info("pokedex API running...")

	// Routes
	pokemon.LandingPage(router, allPokemon)
	pokemon.PokemonById(router, allPokemon)

	router.Run("localhost:8080")

	time.Sleep(5 * time.Second)
}
