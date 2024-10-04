package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// function for getting location_areas from api
func (c *Client) ListLocations(pageURL *string) (RespLocations, error) {
	url := baseURL + "/" + "location-area"

	if pageURL != nil {
		url = *pageURL
	}
	// try request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error with making request")
		return RespLocations{}, err
	}
	// get response
	res, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println("Error getting response")
		return RespLocations{}, err
	}

	defer res.Body.Close()
	// read body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading body")
		return RespLocations{}, err
	}

	locations := RespLocations{}
	// parse json to RespLocations struct
	err = json.Unmarshal(data, &locations)
	if err != nil {
		fmt.Println("Error parsing json")
		return RespLocations{}, err
	}

	return locations, nil
}
