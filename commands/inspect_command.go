package commands

import (
	"fmt"
	"pokedex/config"
)

func inspectCommand(cfg *config.Config, param string) error {
	if param == "" {
		return fmt.Errorf("pokemon parameter is required")
	}

	if val, ok := cfg.Dex.Get(param); ok {
		err := cfg.Dex.Display(val.Name)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
		return nil
	}

	return fmt.Errorf("%s has not been caught", param)
}

var ic = Command{
	name: "inspect",
	desc: "Allow players to see details about a Pokemon if they have caught it before.",
	exec: inspectCommand,
}
