package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("Name is required to inspect")
	}
	name := args[0]

	poke, ok := cfg.caughtPokemon[name]
	if !ok {
		return fmt.Errorf("You have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", poke.Name)
	fmt.Printf("Height: %d\n", poke.Height)
	fmt.Printf("Weight: %d\n", poke.Weight)
	fmt.Println("Stats:")
	for _, stat := range poke.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range poke.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil

}
