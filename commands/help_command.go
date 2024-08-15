package commands

import (
	"fmt"
	"pokedex/config"
)

func helpCommand(cfg *config.Config) error {
	for k, v := range Commands {
		fmt.Println(k, ": ", v.desc)
	}
	return nil
}

var hc = Command{
	name: "help",
	desc: "Displays a help message",
	exec: helpCommand,
}
