package main

import "fmt"

func commandMapf(cfg *config) error {
	locations, err := cfg.client.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locations.Next
	cfg.prevLocationURL = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
