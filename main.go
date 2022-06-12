package main

import (
	pokemon "pokedex-web-app/routes"
	utils "pokedex-web-app/utils"
	data "pokedex-web-app/data"
	

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/styles", "./styles")

	all_pokemon := data.GetPokemonData()

	// Writing logs to file
	utils.LoggingOutput()

	log.Info("Pokedex API running...")

	// Routes
	pokemon.AllPokemon(router, all_pokemon)
	pokemon.PokemonById(router, all_pokemon)

	router.Run("localhost:8000")
}
