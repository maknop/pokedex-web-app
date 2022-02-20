package pokemon

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	views "pokedex-web-app/views"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

func GetAllPokemon(router *gin.Engine) {
	var all_pokemon_data gjson.Result
	pokemon_name := []string{}
	//pokemon_name := make([]string, 152)

	router.GET("/", func(c *gin.Context) {
		log.Info("Request for all pokemon data made.")

		for i := 1; i <= 151; i++ {
			wg := sync.WaitGroup{}
			wg.Add(1)
			go func(idx int) {
				resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", idx))
				if err != nil {
					log.Fatal(err)
				}

				pokemon_data, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}

				all_pokemon_data := gjson.Get(string(pokemon_data), "sprites.front_default")
				for _, name := range all_pokemon_data.Array() {
					pokemon_name = append(pokemon_name, name.String())
					//pokemon_name[idx] = name.String()
				}
				log.Info(all_pokemon_data)

				defer wg.Done()
			}(i)
		}

		views.ViewAllPokemon(c, pokemon_name, all_pokemon_data)

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
