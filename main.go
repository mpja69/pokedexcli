package main

import (
	"time"

	"github.com/mpja69/pokedexcli/internal/pokeapi"
)

type config struct {
	Client   pokeapi.Client
	Next     *string
	Previous *string
	Pokedex  map[string]Pokemon
}
type Pokemon struct {
	Name   string
	Number int
}

func main() {
	cfg := config{
		Client:  pokeapi.NewClient(time.Hour),
		Pokedex: make(map[string]Pokemon),
	}
	startRepl(&cfg)
}
