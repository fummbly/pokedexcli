package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(name string) (Location, error) {
	url := baseURL + "/" + "location-area/" + name

	if val, ok := c.cache.Get(url); ok {
		loc := Location{}
		err := json.Unmarshal(val, &loc)
		if err != nil {
			fmt.Println("Error parsing data from json")
			return Location{}, err
		}
		return loc, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error making request")
		return Location{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println("Error getting response")
		return Location{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading body")
		return Location{}, nil
	}

	loc := Location{}
	err = json.Unmarshal(data, &loc)
	if err != nil {
		fmt.Printf("Error parsing data to json: %+v", err)
		return Location{}, nil
	}

	c.cache.Add(url, data)
	return loc, nil
}
