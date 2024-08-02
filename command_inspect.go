package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Missing Pokemon name")
	}
	name := args[0]
	pokemon, ok := cfg.Pokedex[name]
	if !ok {
		return fmt.Errorf("You haven't cought %s yet.", name)
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats: ")
	for _, val := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", val.Stat.Name, val.BaseStat)
	}
	fmt.Println("Types: ")
	for _, val := range pokemon.Types {
		fmt.Printf(" - %s\n", val.Type.Name)
	}
	return nil
}
