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

	pm := handler.PokeMap{
		Config: config.Config{
			Next:     nil,
			Previous: nil,
		},
		Results: nil,
	}

	var dat []byte
	if val, ok := cfg.Cache.Get(*cfg.Previous); ok {
		dat = val
	} else {
		body, err := helper.GetBody(*cfg.Previous)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
		dat = body
		cfg.Cache.Add(*cfg.Previous, dat)
	}

	err := helper.UnmarshalMap(dat, &pm)
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
