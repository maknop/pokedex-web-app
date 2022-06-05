package pokemon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	pokemon "pokedex-web-app/types"
	views "pokedex-web-app/views"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

func GetAllPokemon(router *gin.Engine) {
	var curr_pokemon pokemon.Pokemon
	var all_pokemon []pokemon.Pokemon

	router.GET("/", func(c *gin.Context) {
		log.Info("Request for all pokemon data made.")

		for i := 1; i <= 151; i++ {
			resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", i))
			if err != nil {
				log.Fatal(err)
			}

			pokemon_data, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			if err := json.NewDecoder(bytes.NewReader(pokemon_data)).Decode(&curr_pokemon); err != nil {
				log.Fatalf("Error parsing json data: %s", err)
			} else {
				log.Println("Pokemon data received")
			}

			fmt.Printf("Pokemon ID: %3d, Pokemon Name: %s\n", curr_pokemon.ID, curr_pokemon.Name)

			all_pokemon = append(all_pokemon, curr_pokemon)
		}

		fmt.Println(all_pokemon)
		views.ViewAllPokemon(c, all_pokemon)

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

		views.ViewPokemonById(c, pokemon_name.String())
	})
}
