package commands

import (
	"fmt"
	"pokedex/config"
	"pokedex/handler"
	"pokedex/helper"
)

func mapCommand(cfg *config.Config, param string) error {
	pm := handler.PokeMap{
		Config: config.Config{
			Next:     nil,
			Previous: nil,
		},
		Results: nil,
	}

	if cfg.Next == nil {
		cfg.Next = new(string)
		*cfg.Next = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	}

	var dat []byte
	if val, ok := cfg.Cache.Get(*cfg.Next); ok {
		dat = val
	} else {
		body, err := helper.GetBody(*cfg.Next)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
		dat = body
		cfg.Cache.Add(*cfg.Next, dat)
	}

	err := helper.UnmarshalMap(dat, &pm)
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
