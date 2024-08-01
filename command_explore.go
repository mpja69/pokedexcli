package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No location given")
	}
	name := args[0]
	res, err := cfg.Client.ListPokemons(name)
	if err != nil {
		return err
	}
	fmt.Println("Pokemons:")
	for _, encounters := range res.PokemonEncounters {
		fmt.Printf(" - %s\n", encounters.Pokemon.Name)
	}

	return nil
}
