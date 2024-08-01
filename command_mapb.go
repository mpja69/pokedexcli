package main

import (
	"errors"
	"fmt"
)

func callbackMapB(cfg *config) error {
	if cfg.Previous == nil {
		return errors.New("Already on the first page!")
	}
	res, err := cfg.Client.ListLocationAreas(cfg.Previous)
	if err != nil {
		return err
	}
	fmt.Println("Locations:")
	for _, loc := range res.Results {
		fmt.Printf(" - %s\n", loc.Name)
	}

	cfg.Next = res.Next
	cfg.Previous = res.Previous
	return nil
}
