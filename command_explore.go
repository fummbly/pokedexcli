package main

import "fmt"

func commandExplore(cfg *config, commandArg string) error {
	if commandArg == "" {
		return fmt.Errorf("Missing location name or id")
	}

	fmt.Printf("Exploring %s...\n", commandArg)
	location, err := cfg.client.GetLocation(commandArg)
	if err != nil {
		return err
	}
	fmt.Println("Found pokemon:")
	for _, pokemon := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
