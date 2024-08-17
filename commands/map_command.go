package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedex/config"
	"pokedex/handler"
	"pokedex/pokecache"
	"time"
)

//func getMap(route string) {
//
//}

func mapCommand(cfg *config.Config) error {
	pm := handler.PokeMap{
		Count: 0,
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

	if cfg.Cache == nil {
		cfg.Cache = new(pokecache.Cache)
		*cfg.Cache = pokecache.NewCache(20 * time.Second)
	}

	// check if data is in the cache
	if val, ok := cfg.Cache.Get(*cfg.Next); ok {
		err := UnmarshalMap(val, &pm)
		fmt.Println("Cache hit")
		if err != nil {
			return fmt.Errorf(err.Error())
		}
	} else {
		// Get request
		res, err := http.Get(*cfg.Next)
		if err != nil {
			return fmt.Errorf(err.Error())
		}

		// Read the response body
		body, err := io.ReadAll(res.Body)
		err = res.Body.Close()

		// throw if the status code is not within the range of 200
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

		cfg.Cache.Add(*cfg.Next, body)
	}

	// check if next if null
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

func UnmarshalMap(body []byte, pm *handler.PokeMap) error {
	// unmarshal the JSON response body into a struct
	err := json.Unmarshal(body, &pm)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	// print the maps
	for _, v := range pm.Results {
		fmt.Println(v.Name)
	}

	return nil
}

var mc = Command{
	name: "map",
	desc: "Displays the names of the next 20 location areas in the Pokemon world",
	exec: mapCommand,
}
