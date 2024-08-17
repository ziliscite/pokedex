package main

import (
	"bufio"
	"os"
	"pokedex/commands"
	"pokedex/config"
	"pokedex/player_pokedex"
	"pokedex/pokecache"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	options := commands.Commands

	client := config.Client{
		Endpoint: "https://pokeapi.co/api/v2/",
		Cache:    pokecache.NewCache(60 * time.Second),
		Dex:      player_pokedex.NewUserPokedex(),
	}

	cfg := &config.Config{
		Client:   client,
		Next:     nil,
		Previous: nil,
	}

	cfg.Next = new(string)
	*cfg.Next = cfg.Endpoint + "location-area/?offset=0&limit=20"

	repl(scanner, options, cfg)
}
