package main

import (
	pokemon "pokedex-api/routes/pokemon"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/css", "./css")

	// Writing logs to file
	loggingOutput()

	log.Info("Pokedex API running...")

	// Routes
	pokemon.GetAllPokemon(router)
	pokemon.GetPokemonById(router)

	router.Run("localhost:8080")
}
