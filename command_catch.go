package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("Pokemon name is required to catch")
	}
	name := args[0]
	poke, err := cfg.client.GetPokemon(name)
	if err != nil {
		return fmt.Errorf("Error getting pokemon info: %v\n", err)
	}
	fmt.Printf("Catching pokemon: %s\n", poke.Name)

	chance := rand.Intn(poke.BaseExperience)

	fmt.Printf("Throwing pokeball at %s...\n", poke.Name)

	if chance > 40 {
		fmt.Printf("Failed to catch: %s\n", poke.Name)
		return nil
	}

	fmt.Printf("Caught %s\n", poke.Name)
	cfg.caughtPokemon[poke.Name] = poke

	return nil

}
