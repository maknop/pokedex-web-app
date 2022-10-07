package main

import (
	"os"
	data "pokedex-web-app/data"
	pokemon "pokedex-web-app/routes"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/styles", "./styles")

	all_pokemon := data.GetPokemonData()

	log.Info("Pokedex API running...")

	// Routes
	pokemon.AllPokemon(router, all_pokemon)
	pokemon.PokemonById(router, all_pokemon)

	router.Run("localhost:8080")
}
