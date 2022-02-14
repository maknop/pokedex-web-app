package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pokemon struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var pokemon = []Pokemon{}

func getPokemon(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pokemon)
}

func main() {
	router := gin.New()

	router.Use(gin.Logger())

	router.GET("/pokemon", func(c *gin.Context) {
		resp, err := http.Get("https://pokeapi.co/api/v2/pokemon")
		if err != nil {
			log.Fatalln(err)
		}

		all_pokemon, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(all_pokemon))

		defer resp.Body.Close()
	})

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

	router.Run("localhost:8080")
}
