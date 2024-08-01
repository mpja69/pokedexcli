package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemons(name string) (PokemonResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint + "/" + name

	data, ok := c.cache.Get(fullURL)
	if ok {
		pokeMap := PokemonResponse{}
		err := json.Unmarshal(data, &pokeMap)
		if err != nil {
			return PokemonResponse{}, err
		}
		return pokeMap, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonResponse{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return PokemonResponse{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	pokeMap := PokemonResponse{}
	err = json.Unmarshal(data, &pokeMap)
	if err != nil {
		return PokemonResponse{}, err
	}
	c.cache.Add(fullURL, data)

	return pokeMap, nil
}
