package pokemon

import (
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func GetAllPokemon(router *gin.Engine) {
	router.GET("/pokemon", func(c *gin.Context) {
		log.Info("Request for all pokemon data made.")
		resp, err := http.Get("https://pokeapi.co/api/v2/pokemon")
		if err != nil {
			log.Fatal(err)
		}

		all_pokemon, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(all_pokemon))

		c.HTML(http.StatusOK, "pokedex.tmpl", gin.H{
			"title": "Main website",
		})

		defer resp.Body.Close()
	})
}

func GetPokemonById(router *gin.Engine) {
	router.GET("/pokemon/:id", func(c *gin.Context) {
		pokemon_id := c.Param("id")
		log.Info("Request for pokemon data with ID ", pokemon_id, " made." )
		resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon_id))

		if err != nil {
			log.Fatal(err)
		}

		all_pokemon, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(all_pokemon))

		c.HTML(http.StatusOK, "pokedex.tmpl", gin.H{
			"title": "Main website",
		})

		defer resp.Body.Close()
	})
}
