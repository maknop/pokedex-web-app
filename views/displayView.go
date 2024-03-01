package views

import (
	"net/http"

	"github.com/gin-gonic/gin"

	pokemon "pokedex-web-app/types"
)

// Endpoint: "/"
func ViewAllPokemon(c *gin.Context, pokemon_data []pokemon.Pokemon) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":       "Pokédex",
		"caught":      "898",
		"seen":        "898",
		"pokemonData": pokemon_data,
	})
}

// Endpoint: "pokemon/:id"
func ViewPokemonById(c *gin.Context, pokemon_id string, pokemon_data pokemon.Pokemon) {
	c.HTML(http.StatusOK, "singleEntry.html", gin.H{
		"pokemonID":   pokemon_id,
		"pokemonData": pokemon_data,
	})
}
