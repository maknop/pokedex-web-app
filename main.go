package main

import (
	api "pokedex-web-app/api"
	routes "pokedex-web-app/routes"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
	router.Static("/styles", "./styles")

	allPokemon := api.GetPokemonData()

	// Routes
	routes.LandingPage(router, allPokemon)
	routes.PokemonById(router, allPokemon)

	router.Run("localhost:8080")

	time.Sleep(5 * time.Second)
}
