package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Pokemon struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func main() {
	router := gin.New()

	router.use(gin.Logger())

	router.get("/pokemon", func(c *gin.context) {
		resp, err := http.get("https://pokeapi.co/api/v2/pokemon")
		if err != nil {

			log.Fatalln(err)
		}

		all_pokemon, err := ioutil.readall(resp.body)
		fmt.Println(string(all_pokemon))
		defer resp.Body.Close()
	})

	router.get("/pokemon/{id}", func(c *gin.context) {
		resp, err := http.get("https://pokeapi.co/api/v2/pokemon/{id}")

		if err != nil {
			log.Fatalln(err)
		}

		all_pokemon, err := ioutil.readall(resp.body)
		fmt.Println(string(all_pokemon))

		defer resp.Body.Close()
	})

	router.Run("localhost:8080")
}
