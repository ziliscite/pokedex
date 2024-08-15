package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedex/config"
	"pokedex/handler"
)

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
		// Allocate memory for the string
		cfg.Next = new(string)
		*cfg.Next = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	}

	res, err := http.Get(*cfg.Next)
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

	err = json.Unmarshal(body, &pm)
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

	for _, v := range pm.Results {
		fmt.Println(v.Name)
	}

	return nil
}

var mc = Command{
	name: "map",
	desc: "Displays a help message",
	exec: mapCommand,
}
