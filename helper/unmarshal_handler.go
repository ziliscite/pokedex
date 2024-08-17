package helper

import (
	"encoding/json"
	"fmt"
	"pokedex/handler"
)

func UnmarshalMap(body []byte, pm *handler.PokeMap) error {
	// unmarshal the JSON response body into a struct
	err := json.Unmarshal(body, &pm)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	for _, v := range pm.Results {
		fmt.Println(v.Name)
	}

	return nil
}

// Generics time?
// probably not

func UnmarshalExplore(body []byte, pe *handler.PokeExplore) error {
	err := json.Unmarshal(body, &pe)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	fmt.Println("Found Pokemon:")
	for _, v := range pe.PokemonEncounters {
		fmt.Println(" - ", v.Name)
	}

	return nil
}
