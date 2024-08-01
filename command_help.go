package main

import "fmt"

func callbackHelp() error {
	commands := getCommands()

	fmt.Println("Available commands:")
	for _, cmd := range commands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}
