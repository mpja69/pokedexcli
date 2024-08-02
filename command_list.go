package main

import "fmt"

func callbackList(cfg *config, name ...string) error {

	fmt.Println("Pokedex:")
	for key, val := range cfg.Pokedex {
		fmt.Printf(" - %s: %d\n", key, val.Number)
	}
	fmt.Println("")
	return nil
}
