package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"

	types "pokedex-web-app/types"
)

func GetPokemonData() []types.Pokemon {

	var currPokemon types.Pokemon
	var allPokemon []types.Pokemon

	log.Info("Request for all pokemon data made.")

	for i := 1; i <= 1025; i++ {
		resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", i))
		if err != nil {
			log.Fatal(err)
		}

		pokemonData, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		if err := json.NewDecoder(bytes.NewReader(pokemonData)).Decode(&currPokemon); err != nil {
			log.Fatalf("Error parsing json data: %s", err)
		} else {
			log.Printf("Pokemon data received - Pokemon ID: %3d, Pokemon Name: %s\n", currPokemon.ID, currPokemon.Name)
		}

		allPokemon = append(allPokemon, currPokemon)
	}

	return allPokemon
}
