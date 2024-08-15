package main

import (
	"bufio"
	"fmt"
	"pokedex/commands"
	"pokedex/config"
)

func repl(scanner *bufio.Scanner, options map[string]commands.Command, cfg *config.Config) {
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		command := formatInput(scanner.Text())
		if _, ok := options[command]; !ok {
			fmt.Printf("Unknown command: %s\n", command)
			continue
		}
		c := options[command]
		err := c.Execute(cfg)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
