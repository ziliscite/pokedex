package main

import (
	"bufio"
	"os"
	"pokedex/commands"
	"pokedex/config"
	"strings"
)

func formatInput(i string) string {
	return strings.ToLower(strings.TrimSpace(i))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	options := commands.Commands
	cfg := &config.Config{
		Cache:    nil,
		Next:     nil,
		Previous: nil,
	}
	repl(scanner, options, cfg)
}
