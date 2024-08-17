package config

import "pokedex/pokecache"

type Config struct {
	Cache    *pokecache.Cache
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
}
