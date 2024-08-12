package commands

import "fmt"

func helpCommand() error {
	for k, v := range Commands {
		fmt.Println(k, ": ", v.desc)
	}
	return nil
}

var hc = Command{
	name: "help",
	desc: "Displays a help message",
	exec: helpCommand,
}
