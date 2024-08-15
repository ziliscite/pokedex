package handler

import "pokedex/config"

type mapResults struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokeMap struct {
	Count int `json:"count"`
	config.Config
	Results []mapResults
}

//api: "https://pokeapi.co/api/v2/location-area/"
