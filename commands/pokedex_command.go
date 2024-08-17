package commands

import (
	"fmt"
	"pokedex/config"
)

func pokedexCommand(cfg *config.Config, param string) error {
	if cfg.Dex.GetLen() == 0 {
		return fmt.Errorf("pokedex is empty")
	}

	cfg.Dex.GetEntries()
	return nil
}

var pc = Command{
	name: "inspect",
	desc: "Allow players to see details about a Pokemon if they have caught it before.",
	exec: pokedexCommand,
}
