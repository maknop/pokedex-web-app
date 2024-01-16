package views

import (
	"net/http"

	"github.com/gin-gonic/gin"

	types "pokedex-web-app/types"
)

// Endpoint: "/"
func ViewAllPokemon(c *gin.Context, pokemonData []types.Pokemon) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":       "Pokédex",
		"caught":      "898",
		"seen":        "898",
		"pokemonData": pokemonData,
	})
}

// Endpoint: "pokemon/:id"
func ViewPokemonById(c *gin.Context, pokemonID string, pokemonData types.Pokemon) {
	c.HTML(http.StatusOK, "singleEntry.tmpl", gin.H{
		"pokemonID":   pokemonID,
		"pokemonData": pokemonData,
	})
}
