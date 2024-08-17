package config

import (
	"pokedex/player_pokedex"
	"pokedex/pokecache"
)

type Client struct {
	Endpoint string
	Cache    pokecache.Cache
	Dex      player_pokedex.UserPokedex
}

type Config struct {
	Client
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
}
