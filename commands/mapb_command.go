package commands

import (
	"fmt"
	"pokedex/config"
	"pokedex/handler"
	"pokedex/helper"
)

func mapbCommand(cfg *config.Config, param string) error {
	if cfg.Previous == nil {
		return fmt.Errorf("at the start of the map")
	}

	var pm = handler.PokeMap{}
	err := helper.GetMap(cfg, &pm, *cfg.Previous)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	cfg.Next = new(string)
	*cfg.Next = *pm.Next

	if pm.Previous != nil {
		cfg.Previous = new(string)
		*cfg.Previous = *pm.Previous
	} else {
		cfg.Previous = nil
	}

	return nil
}

var mbc = Command{
	name: "mapb",
	desc: "Display the names of the previous 20 location areas in the Pokemon world",
	exec: mapbCommand,
}
