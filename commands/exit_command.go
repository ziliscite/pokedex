package commands

import (
	"os"
)

func exitCommand() error {
	os.Exit(0)
	return nil
}

var ec = Command{
	name: "exit",
	desc: "Exit the Pokedex",
	exec: exitCommand,
}
