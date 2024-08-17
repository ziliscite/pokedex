package commands

import (
	"fmt"
	"pokedex/config"
	"pokedex/handler"
	"pokedex/helper"
)

func exploreCommand(cfg *config.Config, param string) error {
	if param == "" {
		return fmt.Errorf("area parameter is required")
	}
	route := cfg.Endpoint + "location-area/" + param

	var dat []byte
	if val, ok := cfg.Cache.Get(param); ok {
		dat = val
	} else {
		body, err := helper.GetBody(route)
		if err != nil {
			return fmt.Errorf(err.Error())
		}

		dat = body
		cfg.Cache.Add(param, dat)
	}

	// why not just declare it once in global scope? immutability or something, idk
	var pe = handler.PokeExplore{}
	err := helper.UnmarshalExplore(dat, &pe)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

var exc = Command{
	name: "explore",
	desc: "Display the name of all the Pokémon in a given area",
	exec: exploreCommand,
}
