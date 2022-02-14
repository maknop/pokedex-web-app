package pokemon

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllPokemon(router *gin.Engine) {
	router.GET("/pokemon", func(c *gin.Context) {
		resp, err := http.Get("https://pokeapi.co/api/v2/pokemon")
		if err != nil {
			log.Fatalln(err)
		}

		all_pokemon, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(all_pokemon))

		defer resp.Body.Close()
	})
}

func getPokemonById(router *gin.Engine) {
	router.GET("/pokemon/:id", func(c *gin.Context) {
		pokemon_id := c.Param("id")
		resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon_id))

		if err != nil {
			log.Fatalln(err)
		}

		all_pokemon, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(all_pokemon))

		defer resp.Body.Close()
	})
}
