package player_pokedex

import "sync"

type Pokemon struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
}

type UserPokedex struct {
	UserExperience int
	Pokedex        map[string]Pokemon
	mu             *sync.Mutex
}

func NewUserPokedex() UserPokedex {
	return UserPokedex{
		UserExperience: 50.0,
		Pokedex:        make(map[string]Pokemon),
		mu:             &sync.Mutex{},
	}
}

func (p *UserPokedex) Get(key string) Pokemon {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.Pokedex[key]
}

func (p *UserPokedex) Set(key string, value Pokemon) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Pokedex[key] = value
}
