package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]

		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid command! (Try 'help')")
			continue
		}
		var args = []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Gets the next 20 locations",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Gets the previous 20 locations",
			callback:    callbackMapB,
		},
		"explore": {
			name:        "explore (location)",
			description: "Gets the Pokemons at the locations",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch (pokemon)",
			description: "Tries to catch the Pokemons by name, based on its XP",
			callback:    callbackCatch,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows the Pokedex list",
			callback:    callbackPokedex,
		},
		"inspect": {
			name:        "inspect (pokemon)",
			description: "Show a Pokemon",
			callback:    callbackInspect,
		},
		"help": {
			name:        "help",
			description: "Prints help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits Pokedex",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
