package main

import (
	"bufio"
	"fmt"
	"os"
	com "pokedex/commands"
	"strings"
)

func formatInput(i string) string {
	return strings.ToLower(strings.TrimSpace(i))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	options := com.Commands
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		command := formatInput(scanner.Text())
		if _, ok := options[command]; !ok {
			fmt.Printf("Unknown command: %s\n", command)
			continue
		}
		c := options[command]
		c.Execute()
	}
}
