package commands

import (
	"fmt"
	"pokedex/config"
	"pokedex/handler"
	"pokedex/helper"
)

func mapCommand(cfg *config.Config, param string) error {
	var pm = handler.PokeMap{}
	err := helper.GetMap(cfg, &pm, *cfg.Next)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	if pm.Next == nil {
		return fmt.Errorf("at the end of the map")
	}

	cfg.Next = new(string)
	*cfg.Next = *pm.Next

	if pm.Previous != nil {
		cfg.Previous = new(string)
		*cfg.Previous = *pm.Previous
	}

	return nil
}

var mc = Command{
	name: "map",
	desc: "Display the names of the next 20 location areas in the Pokemon world",
	exec: mapCommand,
}
