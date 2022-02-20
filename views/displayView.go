package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

// Endpoint: "/"
func ViewAllPokemon(c *gin.Context, pokemon_name []string, all_pokemon_data gjson.Result) {
	c.HTML(http.StatusOK, "allPokemon.tmpl.html", gin.H{
		"title":         "Pokédex",
		"caught":        "1118",
		"seen":          "1118",
		"pokemonImages": pokemon_name,
		"countPokemon":  all_pokemon_data,
	})
}

// Endpoint: "pokemon/:id"
func ViewPokemonById(c *gin.Context, pokemon_name string) {
	c.HTML(http.StatusOK, "singlePokemon.tmpl.html", gin.H{
		"pokemonName": pokemon_name,
	})
}
