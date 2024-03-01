package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	log "github.com/sirupsen/logrus"

	pokemon "pokedex-web-app/types"
)

func GetPokemonData() []pokemon.Pokemon {
	var currPokemon pokemon.Pokemon
	var allPokemon []pokemon.Pokemon

	wg := sync.WaitGroup{}
	var m sync.Mutex

	log.Info("Request for all pokemon data made.")

	for i := 1; i <= 25; i++ {
		wg.Add(1)

		go func(i int) {
			m.Lock()

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

			m.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()

	return allPokemon
}
