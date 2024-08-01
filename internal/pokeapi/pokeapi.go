package pokeapi

import (
	"net/http"
	"time"

	"github.com/mpja69/pokedexcli/internal/pokecash"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
	cache      pokecash.Cache
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecash.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
