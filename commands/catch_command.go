package commands

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"pokedex/config"
	"pokedex/helper"
	"pokedex/player_pokedex"
	"time"
)

func catchCommand(cfg *config.Config, param string) error {
	if param == "" {
		return fmt.Errorf("pokemon parameter is required")
	}

	if val, ok := cfg.Dex.Get(param); ok {
		fmt.Printf("%s is already caught!\n", val.Name)
		return nil
	}

	route := cfg.Endpoint + "pokemon/" + param

	// After deliberate consideration, I don't want to wait for each unsuccessful attempt
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

	pokemon := player_pokedex.Pokemon{}
	err := json.Unmarshal(dat, &pokemon)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	ticker := time.NewTicker(500 * time.Millisecond)
	fmt.Printf("Throwing a Pokeball at %s", param)
	go func() {
		for range ticker.C {
			fmt.Print(".")
		}
	}()
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println()

	res := rand.Intn(pokemon.BaseExperience)
	if res > cfg.Dex.UserExperience {
		fmt.Println(param, "was caught!")
		cfg.Dex.Set(param, pokemon)
		return nil
	}
	fmt.Println(param, "has escaped...")
	return nil
}

var cc = Command{
	name: "catch",
	desc: "Catch a Pokemon and adds them to the user's Pokedex",
	exec: catchCommand,
}
