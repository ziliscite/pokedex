package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedex/config"
	"pokedex/handler"
)

func bmapCommand(cfg *config.Config) error {
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

	err = json.Unmarshal(body, &pm)
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

	for _, v := range pm.Results {
		fmt.Println(v.Name)
	}

	return nil
}

var bmc = Command{
	name: "map",
	desc: "Displays a help message",
	exec: bmapCommand,
}
