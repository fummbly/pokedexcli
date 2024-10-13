package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fummbly/pokedexcli/internal/pokeapi"
)

// config struct for commands
type config struct {
	client          pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
	caughtPokemon   map[string]pokeapi.Pokemon
}

// startRepl function for starting a repl session
func startRepl(cfg *config) {
	// scanner for scanning stdin
	scanner := bufio.NewScanner(os.Stdin)
	//infinite loop
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()

		// clean input by setting input to lower and seperating words
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		// taking the first word of the input as a command
		commandName := words[0]
		args := []string{}

		if len(words) > 1 {
			args = words[1:]

		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)

			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

// cleans input by setting to lower and seperating words
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

// command struct with name description and callback function
type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

// function for defining commands
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Prints the next 20 locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Prints the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Prints the pokemon found in a given area - use: explore <area>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a given pokemon - use: catch <pokemon>",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspects a caught pokemon and prints its stats",
			callback:    commandInspect,
		},
	}
}
