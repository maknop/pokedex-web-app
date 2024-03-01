package handler

import (
	pokemon "pokedex-web-app/types"
	views "pokedex-web-app/views"

	"github.com/gin-gonic/gin"
	//log "github.com/sirupsen/logrus"
)

func AllPokemon(router *gin.Engine, all_pokemon []pokemon.Pokemon) {
	router.GET("/", func(c *gin.Context) {
		views.ViewAllPokemon(c, all_pokemon)
	})
}

func PokemonById(router *gin.Engine, all_pokemon []pokemon.Pokemon) {
	var curr_pokemon pokemon.Pokemon

	router.GET("/pokemon/:id", func(c *gin.Context) {
		pokemon_id := c.Param("id")

		for _, single := range all_pokemon {
			if single.Name == pokemon_id {
				curr_pokemon = single
			}
		}

		views.ViewPokemonById(c, pokemon_id, curr_pokemon)
	})
}
