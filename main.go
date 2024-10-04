package main

import (
	"time"

	"github.com/fummbly/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		client: pokeClient,
	}

	startRepl(cfg)

}
