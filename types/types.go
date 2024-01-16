package types

type Pokemon struct {
	ID     int           `json:"id"`
	Name   string        `json:"name"`
	Height int           `json:"height"`
	Weight int           `json:"weight"`
	Images PokemonImages `json:"Sprites"`
}

type PokemonImages struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
