package main

import "fmt"

// map forward function
func commandMapf(cfg *config, commandArg string) error {

	// get the locations from the pokeapi
	locations, err := cfg.client.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}
	// set the config next/prev the next/prev in the locations
	cfg.nextLocationURL = locations.Next
	cfg.prevLocationURL = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config, commandArg string) error {
	locations, err := cfg.client.ListLocations(cfg.prevLocationURL)
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
