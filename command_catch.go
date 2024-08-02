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
	res, err := cfg.Client.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Cathing %s...", name)
	xp := res.BaseExperience
	rnd := rand.Intn(400)
	if rnd > xp {
		pokemon, ok := cfg.Pokedex[name]
		nbr := 0
		if ok {
			nbr = pokemon.Number
		}
		// HACK: Ska jag bara uppdatera antalet??
		cfg.Pokedex[name] = Pokemon{Name: name, Number: nbr + 1}
		fmt.Println("Got'em!")
	} else {
		fmt.Println("Missed!")
	}
	return nil
}
