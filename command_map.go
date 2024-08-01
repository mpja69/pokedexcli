package main

import (
	"fmt"
)

func callbackMap(cfg *config) error {
	res, err := cfg.Client.ListLocationAreas(cfg.Next)
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
