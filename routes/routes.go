package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const URL = "https://beta.pokeapi.co/graphql/v1beta"

func GetPokemon() {
	// Query
	// query := `
	// 	{
	// 		query samplePokeAPIquery {
	// 			pokemon_v2_pokemonsprites {
	// 			  sprites
	// 			  pokemon_v2_pokemon {
	// 				weight
	// 				height
	// 				id
	// 				name
	// 			  }
	// 		    }
	// 		}
	// 	}
	// `

	query := map[string]string{
		"query": `
			{
				pokemon_v2_pokemonsprites {
			        sprites
			        pokemon_v2_pokemon {
				      weight
				      height
				      id
				      name
			        }
			      }
		        }
		`,
	}

	jsonValue, _ := json.Marshal(query)
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonValue))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "*/*")

	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	defer response.Body.Close()

	data, _ := io.ReadAll(response.Body)
	fmt.Println(string(data))

}
