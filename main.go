package main

import "github.com/mpja69/pokedexcli/internal/pokeapi"

type config struct {
	Client   pokeapi.Client
	Next     *string
	Previous *string
}

func main() {
	cfg := config{
		Client: pokeapi.NewClient(),
	}
	// config := config{"https://pokeapi.co/api/v2/location-area/", "https://pokeapi.co/api/v2/location-area/"}
	startRepl(&cfg)
}
