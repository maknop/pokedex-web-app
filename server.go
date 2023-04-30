package main

import (
	"log"
	data "pokedex-web-app/data"
	pokemon "pokedex-web-app/routes"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/styles", "./styles")

	all_pokemon := data.GetPokemonData()

	log.Println("Pokedex API running")

	// Routes
	pokemon.AllPokemon(router, all_pokemon)
	pokemon.PokemonById(router, all_pokemon)

	router.Run("localhost:8080")

	time.Sleep(5 * time.Second)
}
