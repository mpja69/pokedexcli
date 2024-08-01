package main

import (
	"time"

	"github.com/mpja69/pokedexcli/internal/pokeapi"
)

type config struct {
	Client   pokeapi.Client
	Next     *string
	Previous *string
}

func main() {
	cfg := config{
		Client: pokeapi.NewClient(time.Hour),
	}
	startRepl(&cfg)
}
