package helper

import (
	"fmt"
	"pokedex/config"
	"pokedex/handler"
)

func GetMap(cfg *config.Config, pm *handler.PokeMap, target string) error {
	var dat []byte
	if val, ok := cfg.Cache.Get(target); ok {
		dat = val
	} else {
		body, err := GetBody(target)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
		dat = body
		cfg.Cache.Add(target, dat)
	}

	err := UnmarshalMap(dat, pm)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}
