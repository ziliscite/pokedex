package commands

import (
	"os"
	"pokedex/config"
)

func exitCommand(cfg *config.Config, param string) error {
	os.Exit(0)
	return nil
}

var ec = Command{
	name: "exit",
	desc: "Exit the Pokedex",
	exec: exitCommand,
}
