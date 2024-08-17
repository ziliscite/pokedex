package main

import (
	"bufio"
	"fmt"
	"pokedex/commands"
	"pokedex/config"
	"pokedex/helper"
)

func repl(scanner *bufio.Scanner, options map[string]commands.Command, cfg *config.Config) {
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		command := helper.FormatInput(scanner.Text())
		if _, ok := options[command[0]]; !ok {
			fmt.Printf("Unknown command: %s\n", command)
			continue
		}

		c := options[command[0]]
		err := c.Execute(cfg, command[1])

		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
