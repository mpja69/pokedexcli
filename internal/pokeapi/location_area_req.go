package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	// "github.com/mpja69/pokedexcli/internal/pokeapi"
)

func (c *Client) ListLocationAreas(nextURL *string) (LocationAreaResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if nextURL != nil {
		fullURL = *nextURL
	}
	data, ok := c.cache.Get(fullURL)
	if ok {
		pokeMap := LocationAreaResponse{}
		err := json.Unmarshal(data, &pokeMap)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		return pokeMap, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreaResponse{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	pokeMap := LocationAreaResponse{}
	err = json.Unmarshal(data, &pokeMap)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	c.cache.Add(fullURL, data)

	return pokeMap, nil
}
