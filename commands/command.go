package commands

import (
	"fmt"
	"pokedex/config"
)

type Executable interface {
	Execute() error
}

type Command struct {
	name string
	desc string
	exec func(cfg *config.Config) error
}

func (c *Command) Execute(cfg *config.Config) error {
	err := c.exec(cfg)
	if err != nil {
		return fmt.Errorf("%v: %v", c.name, err)
	}
	return nil
}

var Commands map[string]Command

func init() {
	Commands = map[string]Command{
		"help": hc,
		"exit": ec,
		"map":  mc,
		"bmap": bmc,
	}
}
