package handler

import "pokedex/config"

type mapResults struct {
	Name string `json:"name"`
}

type PokeMap struct {
	config.Config
	Results []mapResults `json:"results"`
}

//api: "https://pokeapi.co/api/v2/location-area/"
