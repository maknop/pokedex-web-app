package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pokemon struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func main() {
	router := gin.New()

	router.Use(gin.Logger())

	router.GET("/pokemon", func(c *gin.Context) {
		resp, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/pokemon", nil)
		if err != nil {

			log.Fatalln(err)
		}

		all_pokemon, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(all_pokemon))
		defer resp.Body.Close()
	})

	router.GET("/pokemon/{id}", func(c *gin.Context) {
		resp, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/pokemon/{id}", nil)

		if err != nil {
			log.Fatalln(err)
		}

		all_pokemon, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(all_pokemon))

		defer resp.Body.Close()
	})

	router.Run("localhost:8080")
}
