package main

import (
	"os"
	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	pokemon "pokedex-api/routes/pokemon"
)

type Pokemon struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func loggingOutput() {
	logFile, _ := os.Create("logs/api_requests.log")

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(logFile)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")

	// Writing logs to file
	loggingOutput()

	log.Info("Pokedex API running.")

	// Routes
	pokemon.GetAllPokemon(router)
	pokemon.GetPokemonById(router)

	router.Run("localhost:8080")
}
