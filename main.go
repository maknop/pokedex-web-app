package main

import (
	data "pokedex-web-app/data"
	pokemon "pokedex-web-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/styles", "./styles")

	all_pokemon := data.GetPokemonData()

	//log.Info("Pokedex API running...")

	// Routes
	pokemon.AllPokemon(router, all_pokemon)
	pokemon.PokemonById(router, all_pokemon)

	router.Run()
}
