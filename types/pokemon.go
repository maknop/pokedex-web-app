package pokemon

type Pokemon struct {
	ID     int           `json:"id"`
	Name   string        `json:"name"`
	Height int           `json:"height"`
	Weight int           `json:"weight"`
	Types  PokemonType   `json:"-"`
	Images PokemonImages `json:"Sprites"`
}

type PokemonType struct {
}

type PokemonImages struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
