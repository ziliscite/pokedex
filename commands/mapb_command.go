package commands

import (
	"fmt"
	"io"
	"net/http"
	"pokedex/config"
	"pokedex/handler"
	"pokedex/pokecache"
	"time"
)

func mapbCommand(cfg *config.Config) error {
	if cfg.Previous == nil {
		return fmt.Errorf("at the start of the map")
	}

	pm := handler.PokeMap{
		Count: 0,
		Config: config.Config{
			Next:     nil,
			Previous: nil,
		},
		Results: nil,
	}

	if cfg.Cache == nil {
		cfg.Cache = new(pokecache.Cache)
		*cfg.Cache = pokecache.NewCache(20 * time.Second)
	}

	if val, ok := cfg.Cache.Get(*cfg.Previous); ok {
		err := UnmarshalMap(val, &pm)
		fmt.Println("Cache hit")
		if err != nil {
			return fmt.Errorf(err.Error())
		}
	} else {
		res, err := http.Get(*cfg.Previous)
		if err != nil {
			return fmt.Errorf(err.Error())
		}

		body, err := io.ReadAll(res.Body)
		err = res.Body.Close()
		if res.StatusCode > 299 {
			return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			return fmt.Errorf(err.Error())
		}

		err = UnmarshalMap(body, &pm)
		if err != nil {
			return fmt.Errorf(err.Error())
		}

		cfg.Cache.Add(*cfg.Previous, body)
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
	desc: "Displays the names of the previous 20 location areas in the Pokemon world",
	exec: mapbCommand,
}
