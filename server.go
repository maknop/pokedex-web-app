package main

import (
	"pokedex-api/routes/pokemonHandler"

	"github.com/gin-gonic/gin"
)

type Pokemon struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	router := gin.Default()

	router.Use(gin.Logger())

	pokemonHandler.Routes(router)
	pokemonHandler.Routes(router)

	router.Run("localhost:8080")
}
