package main

import "fmt"

func commandExplore(cfg *config, commandArg string) error {
	if commandArg == "" {
		return fmt.Errorf("Missing location name or id")
	}

	location, err := cfg.client.GetLocation(commandArg)
	if err != nil {
		return err
	}

	for _, pokemon := range location.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
