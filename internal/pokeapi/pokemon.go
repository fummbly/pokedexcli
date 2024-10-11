package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemon string) (Pokemon, error) {

	if pokemon == "" {
		return Pokemon{}, fmt.Errorf("Pokemon required to make request")
	}

	url := baseURL + "/pokemon/" + pokemon

	if value, ok := c.cache.Get(url); ok {
		poke := Pokemon{}
		err := json.Unmarshal(value, &poke)
		if err != nil {
			fmt.Println("Error parsing from json to Pokemon struct")
			return Pokemon{}, err
		}
		return poke, nil

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error making Pokemon get request")
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println("Error getting response from Pokemon request")
		return Pokemon{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading body for Pokemon response")
		return Pokemon{}, err
	}

	poke := Pokemon{}

	err = json.Unmarshal(data, &poke)
	if err != nil {
		fmt.Println("Error parsing data to Pokemon struct")
		return Pokemon{}, err
	}

	c.cache.Add(url, data)
	return poke, nil

}
