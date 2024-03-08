package main

import (
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	data "github.com/maknop/pokedex-web-app/data"
	db "github.com/maknop/pokedex-web-app/database"
	pokemon "github.com/maknop/pokedex-web-app/routes"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
	router.Static("/styles", "./styles")

	db.Connect()

	allPokemon := data.GetPokemonData()

	log.Info("pokedex API running...")

	// Routes
	pokemon.LandingPage(router, allPokemon)
	pokemon.PokemonById(router, allPokemon)

	router.Run("localhost:8080")
}
