package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("You must include an area")
	}

	fmt.Printf("Exploring %s...\n", args[0])
	location, err := cfg.client.GetLocation(args[0])
	if err != nil {
		return err
	}
	fmt.Println("Found pokemon:")
	for _, pokemon := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
