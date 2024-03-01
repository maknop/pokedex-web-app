package handler

import (
	"github.com/gin-gonic/gin"

	types "pokedex-web-app/types"
	views "pokedex-web-app/views"
)

func LandingPage(router *gin.Engine, allPokemon []types.Pokemon) {
	router.GET("/", func(c *gin.Context) {
		views.ViewAllPokemon(c, allPokemon)
	})
}

func PokemonById(router *gin.Engine, allPokemon []types.Pokemon) {
	var currPokemon types.Pokemon

	router.GET("/pokemon/:id", func(c *gin.Context) {
		pokemon_id := c.Param("id")

		for _, single := range allPokemon {
			if single.Name == pokemon_id {
				currPokemon = single
			}
		}

		views.ViewPokemonById(c, pokemon_id, currPokemon)
	})
}
