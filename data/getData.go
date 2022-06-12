package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"

	pokemon "pokedex-web-app/types"
)

func GetPokemonData() []pokemon.Pokemon {
	var curr_pokemon pokemon.Pokemon
	var all_pokemon []pokemon.Pokemon

	log.Info("Request for all pokemon data made.")

	for i := 1; i <= 898; i++ {
		resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", i))
		if err != nil {
			log.Fatal(err)
		}

		pokemon_data, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		if err := json.NewDecoder(bytes.NewReader(pokemon_data)).Decode(&curr_pokemon); err != nil {
			log.Fatalf("Error parsing json data: %s", err)
		} else {
			log.Printf("Pokemon data received - Pokemon ID: %3d, Pokemon Name: %s\n", curr_pokemon.ID, curr_pokemon.Name)
		}

		all_pokemon = append(all_pokemon, curr_pokemon)
	}

	return all_pokemon
}
