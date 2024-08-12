package commands

import "fmt"

type Executable interface {
	Execute() error
}

type Command struct {
	name string
	desc string
	exec func() error
}

func (c *Command) Execute() {
	err := c.exec()
	if err != nil {
		_ = fmt.Errorf("%v: %v", c.name, err)
	}
}

var Commands map[string]Command

func init() {
	Commands = map[string]Command{
		"help": hc,
		"exit": ec,
	}
}
