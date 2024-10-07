package pokeapi

import (
	"testing"
	"time"
)

func TestLocationParsing(t *testing.T) {
	client := NewClient(5*time.Second, 5*time.Minute)

	loc, err := client.GetLocation("canalave-city-area")

	if err != nil {
		t.Errorf("Error getting location: %v", err)
		return
	}

	if loc.ID != 1 {
		t.Errorf("Error not ID expected by call")
		return
	}

	if loc.PokemonEncounters[0].Pokemon.Name != "tentacool" {
		t.Errorf("Error not pokemon expected by call")
		return
	}

}
