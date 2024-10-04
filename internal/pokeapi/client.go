package pokeapi

import (
	"net/http"
	"time"
)

// define the client struct for pokeapi
type Client struct {
	httpClient http.Client
}

// create a new client
func NewClient(timeout time.Duration) Client {
	return Client{
		http.Client{
			Timeout: timeout,
		},
	}
}
