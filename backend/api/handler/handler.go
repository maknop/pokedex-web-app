package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Operation struct {
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
	OperationName string                 `json:"operationName"`
}

const URL = "https://beta.pokeapi.co/graphql/v1beta"

var (
	query = Operation{
		OperationName: "pokemon_details",
		Variables: map[string]interface{}{
			"name": "staryu",
		},
		Query: ``,
	}
)

func allPokemon(w http.ResponseWriter, req *http.Request) {
	log.Info(w, "All the Pokemon!\n")
}

func pokemonDetails(w http.ResponseWriter, req *http.Request) {
	log.Info("requesting details for pokemon []")
	body, err := json.Marshal(query)
	if err != nil {
		log.Error(err)
	}

	resp, err := http.Post(URL, "", bytes.NewReader(body))
	if err != nil {
		log.Error(err)
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}
	log.Info(string(body))
}

func main() {
	log.Info("server running at http://localhost:8080")

	log.Info("serving endpoint at /pokemon")
	http.HandleFunc("/pokemon", allPokemon)

	log.Info("serving endpoint at /pokemon/:name")
	http.HandleFunc("/pokemon/:name", pokemonDetails)

	log.Error(http.ListenAndServe(":8000", nil))
}
