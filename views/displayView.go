package views

import (
	"net/http"

	"github.com/gin-gonic/gin"

	pokemon "pokedex-web-app/types"
)

// Endpoint: "/"
func ViewAllPokemon(c *gin.Context, pokemon_data []pokemon.Pokemon) {
	c.HTML(http.StatusOK, "allPokemon.tmpl.html", gin.H{
		"title":       "Pokédex",
		"caught":      "1118",
		"seen":        "1118",
		"pokemonData": pokemon_data,
	})
}

// Endpoint: "pokemon/:id"
func ViewPokemonById(c *gin.Context, pokemon_name string) {
	c.HTML(http.StatusOK, "singlePokemon.tmpl.html", gin.H{
		"pokemonName": pokemon_name,
	})
}
