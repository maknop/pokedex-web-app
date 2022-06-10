package main

import (
	pokemon "pokedex-web-app/routes"
	utils "pokedex-web-app/utils"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/styles", "./styles")

	// Writing logs to file
	utils.LoggingOutput()

	log.Info("Pokedex API running...")

	// Routes
	pokemon.GetAllPokemon(router)
	pokemon.GetPokemonById(router)

	router.Run("localhost:8000")
}
