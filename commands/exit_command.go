package commands

import (
	"os"
	"pokedex/config"
)

func exitCommand(cfg *config.Config) error {
	os.Exit(0)
	return nil
}

var ec = Command{
	name: "exit",
	desc: "Exit the Pokedex",
	exec: exitCommand,
}
