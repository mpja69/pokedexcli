package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" >")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		command := cleaned[0]
		switch command {
		case "exit":
			os.Exit(0)
		case "help":
			fmt.Println("Available commands:")
			fmt.Println(" - help: This list")
			fmt.Println(" - exit: Exit REPL")
			fmt.Println("")
		default:
			fmt.Println("invalid command")
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func()
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
