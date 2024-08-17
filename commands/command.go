package commands

import (
	"fmt"
	"pokedex/config"
)

type Executable interface {
	Execute(cfg *config.Config) error
}

type Command struct {
	name string
	desc string
	exec func(cfg *config.Config, param string) error
}

func (c *Command) Execute(cfg *config.Config, param string) error {
	err := c.exec(cfg, param)
	if err != nil {
		return fmt.Errorf("%v: %v", c.name, err)
	}
	return nil
}

var Commands map[string]Command

func init() {
	Commands = map[string]Command{
		"help":    hc,
		"exit":    ec,
		"map":     mc,
		"mapb":    mbc,
		"explore": exc,
		"catch":   cc,
		"inspect": ic,
	}
}
