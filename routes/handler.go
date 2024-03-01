package handler

import (
	types "pokedex-web-app/types"
	views "pokedex-web-app/views"

	"github.com/gin-gonic/gin"
	//log "github.com/sirupsen/logrus"
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
