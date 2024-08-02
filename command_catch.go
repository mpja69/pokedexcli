package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Missing Pokemon name")
	}
	name := args[0]
	pokemon, err := cfg.Client.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Cathing %s...", name)
	xp := pokemon.BaseExperience
	rnd := rand.Intn(400)
	if rnd > xp {
		cfg.Pokedex[name] = pokemon
		fmt.Println("Got'em!")
	} else {
		fmt.Println("Missed!")
	}
	return nil
}
