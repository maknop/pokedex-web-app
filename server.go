package main

import (
	pokemon "pokedex-api/routes/pokemon"

	"github.com/gin-gonic/gin"
)

type Pokemon struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	router := gin.Default()

	router.Use(gin.Logger())

	pokemon.GetAllPokemon(router)
	pokemon.GetPokemonById(router)

	router.Run("localhost:8080")
}
