package main

import (
	pokemon "github.com/maknop/pokedex-api/routes/pokemon"

	"github.com/gin-gonic/gin"
)

type Pokemon struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	router := gin.Default()

	router.Use(gin.Logger())

	pokemon.Routes(router)
	pokemon.Routes(router)

	router.Run("localhost:8080")
}
