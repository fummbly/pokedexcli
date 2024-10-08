package main

import (
	"time"

	"github.com/fummbly/pokedexcli/internal/pokeapi"
)

func main() {
	// Creating a pokeCLient for the config
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		client: pokeClient,
	}
	// starting repl session
	startRepl(cfg)

}
