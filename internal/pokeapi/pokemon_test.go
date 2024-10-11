package pokeapi

import (
	"testing"
	"time"
)

func TestPokemonGet(t *testing.T) {

	client := NewClient(5*time.Second, 5*time.Minute)
	poke, err := client.GetPokemon("1")

	if err != nil {
		t.Errorf("Error getting pokemon information: %v", err)
		return
	}

	if poke.ID != 1 {
		t.Errorf("Error expecting different id: %d", poke.ID)
		return
	}

	if poke.Name != "bulbasaur" {
		t.Errorf("Error expecting different pokemon name: %s", poke.Name)
		return
	}

}
