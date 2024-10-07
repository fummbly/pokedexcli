package pokeapi

import (
	"net/http"
	"time"

	"github.com/fummbly/pokedexcli/internal/pokecache"
)

// define the client struct for pokeapi
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// create a new client
func NewClient(timeout, interval time.Duration) Client {
	return Client{
		pokecache.NewCache(interval),
		http.Client{
			Timeout: timeout,
		},
	}
}
