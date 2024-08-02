package main

import (
	"time"

	"github.com/mpja69/pokedexcli/internal/pokeapi"
)

type config struct {
	Client   pokeapi.Client
	Next     *string
	Previous *string
	Pokedex  map[string]pokeapi.PokemonResponse
}

func main() {
	cfg := config{
		Client:  pokeapi.NewClient(time.Hour),
		Pokedex: make(map[string]pokeapi.PokemonResponse),
	}
	startRepl(&cfg)
}
