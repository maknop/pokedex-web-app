package pokemon

type Pokemon struct {
	ID     int           `json:"id"`
	Name   string        `json:"name"`
	Images PokemonImages `json:"Sprites"`
}

type PokemonImages struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
