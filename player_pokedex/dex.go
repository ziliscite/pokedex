package player_pokedex

import (
	"fmt"
	"sync"
)

type pokeStats struct {
	Stat struct {
		Name string `json:"name"`
	} `json:"stat"`
	BaseStat int `json:"base_stat"`
}

type pokeType struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

type Pokemon struct {
	Id             int         `json:"id"`
	Name           string      `json:"name"`
	BaseExperience int         `json:"base_experience"`
	Height         int         `json:"height"`
	Weight         int         `json:"weight"`
	Stats          []pokeStats `json:"stats"`
	Types          []pokeType  `json:"types"`
}

type UserPokedex struct {
	UserExperience int
	pokedex        map[string]Pokemon
	mu             *sync.Mutex
}

func NewUserPokedex() UserPokedex {
	return UserPokedex{
		UserExperience: 50.0,
		pokedex:        make(map[string]Pokemon),
		mu:             &sync.Mutex{},
	}
}

func (p *UserPokedex) Get(key string) (*Pokemon, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	entry, ok := p.pokedex[key]
	if !ok {
		return nil, false
	}

	return &entry, true
}

func (p *UserPokedex) Set(key string, value Pokemon) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.pokedex[key] = value
}

func (p *UserPokedex) Display(key string) error {
	val, ok := p.Get(key)
	if !ok {
		return fmt.Errorf("UserPokedex does not have a pokemon with name %s", key)
	}
	fmt.Println("Name:", val.Name)
	fmt.Println("Height:", val.Height)
	fmt.Println("Weight:", val.Weight)
	fmt.Println("Stats:")
	for _, stat := range val.Stats {
		fmt.Println(" - ", stat.Stat.Name, ":", stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range val.Types {
		fmt.Println(" - ", typ.Type.Name)
	}
	return nil
}
