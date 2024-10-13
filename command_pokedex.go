package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("  - Empty")
		return nil
	}
	for _, poke := range cfg.caughtPokemon {
		fmt.Printf("  - %s\n", poke.Name)
	}
	return nil
}
