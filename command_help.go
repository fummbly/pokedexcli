package main

import "fmt"

// prints the name and description of each command
func commandHelp(cfg *config, args ...string) error {
	if len(args) > 0 {
		command, exists := getCommands()[args[0]]
		if exists {
			fmt.Printf("%s: %s\n", command.name, command.description)
			return nil
		}
	}
	fmt.Println()
	fmt.Println("Welcome to the Pokedex")
	fmt.Println("Usage:")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
