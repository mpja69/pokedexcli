package main

import "fmt"

func callbackPokedex(cfg *config, name ...string) error {

	fmt.Println("Pokedex:")
	for key, _ := range cfg.Pokedex {
		fmt.Printf(" - %s\n", key)
	}
	fmt.Println("")
	return nil
}
