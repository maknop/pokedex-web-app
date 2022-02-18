package pokemon

import (
	"fmt"
	"io"
	"net/http"

	//"bytes"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

func GetAllPokemon(router *gin.Engine) {
	var all_pokemon_data gjson.Result
	pokemon_name := []string{}

	router.GET("/pokemon", func(c *gin.Context) {
		log.Info("Request for all pokemon data made.")
		resp, err := http.Get("https://pokeapi.co/api/v2/pokemon")
		if err != nil {
			log.Fatal(err)
		}

		//// Redirect to other endpoint
		//c.Request.URL.Path = "/pokemon/:id"
		//router.HandleContext(c)

		pokemon_data, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		for i := range pokemon_data {
			all_pokemon_data := gjson.Get(string(pokemon_data), fmt.Sprintf("results.%d.name", i))
			for _, name := range all_pokemon_data.Array() {
				pokemon_name = append(pokemon_name, name.String())
			}
			log.Info(all_pokemon_data)
		}

		log.Info(all_pokemon_data)

		c.HTML(http.StatusOK, "allPokemon.tmpl.html", gin.H{
			"title":        "Pokédex",
			"caught":       "1118",
			"seen":         "1118",
			"allPokemon":   pokemon_name,
			"countPokemon": all_pokemon_data,
		})

	})
}

func GetPokemonById(router *gin.Engine) {
	router.GET("/pokemon/:id", func(c *gin.Context) {
		pokemon_id := c.Param("id")
		log.Info("Request for pokemon data with ID ", pokemon_id, " made.")

		resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon_id))
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		p, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		pokemon_name := gjson.Get(string(p), "name")

		log.Info(pokemon_name)

		c.HTML(http.StatusOK, "singlePokemon.tmpl.html", gin.H{
			"pokemonName": pokemon_name,
		})

	})
}
