package handler

import (
	"github.com/gin-gonic/gin"

	models "pokedex-web-app/models"
	views "pokedex-web-app/views"
)

func LandingPage(router *gin.Engine, allPokemon []models.Pokemon) {
	router.GET("/", func(c *gin.Context) {
		views.ViewAllPokemon(c, allPokemon)
	})
}

func PokemonById(router *gin.Engine, allPokemon []models.Pokemon) {
	var currentPokemon models.Pokemon

	router.GET("/pokemon/:id", func(c *gin.Context) {
		pokemon_id := c.Param("id")

		for _, single := range allPokemon {
			if single.Name == pokemon_id {
				currentPokemon = single
			}
		}

		views.ViewPokemonById(c, pokemon_id, currentPokemon)
	})
}
