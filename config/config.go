package config

import (
	"pokedex/pokecache"
)

type Client struct {
	Endpoint string
	Cache    pokecache.Cache
}

type Config struct {
	Client
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
}
