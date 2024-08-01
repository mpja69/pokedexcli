package main

import (
	"./internal/pokeapi"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func callbackMap(c *config) error {
	res, err := http.Get(c.Next)
	if err != nil {
		fmt.Println("Error in Get")
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		println("Response failed...status code: %d, body: %s", res.StatusCode, body)
	}
	if err != nil {
		fmt.Println("Error on ReadAll")
		return err
	}
	pokeMap := PokeMap{}
	err = json.Unmarshal(body, &pokeMap)
	if err != nil {
		fmt.Println("Error in unmarshal")
		return err
	}
	for _, location := range pokeMap.Results {
		fmt.Println(location.Name)
	}
	c.Next = pokeMap.Next
	return nil
}
