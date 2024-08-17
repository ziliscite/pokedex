package handler

type pokemonName struct {
	Name string `json:"name"`
}

type PokemonEncounters struct {
	pokemonName `json:"pokemon"`
}

type PokeExplore struct {
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}
