package main

import "fmt"

// prints the name and description of each command
func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex")
	fmt.Println("Usage:")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
