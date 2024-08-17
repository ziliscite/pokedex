package main

import (
	"bufio"
	"os"
	"pokedex/commands"
	"pokedex/config"
	"pokedex/pokecache"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	options := commands.Commands

	cfg := &config.Config{
		Cache:    new(pokecache.Cache),
		Next:     nil,
		Previous: nil,
	}

	*cfg.Cache = pokecache.NewCache(60 * time.Second)
	repl(scanner, options, cfg)
}
