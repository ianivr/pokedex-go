package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationPokemon(locationID string) (LocationPokemonResponse, error) {
	url := baseURL + "/location-area/" + locationID + "/"

	if data, ok := c.cache.Get(url); ok {
		var locationPokemonResp LocationPokemonResponse
		err := json.Unmarshal(data, &locationPokemonResp)
		if err != nil {
			return LocationPokemonResponse{}, err
		}
		return locationPokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationPokemonResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationPokemonResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationPokemonResponse{}, err
	}

	if res.StatusCode > 299 {
		return LocationPokemonResponse{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
	}

	var locationPokemonResp LocationPokemonResponse
	err = json.Unmarshal(data, &locationPokemonResp)
	if err != nil {
		return LocationPokemonResponse{}, err
	}

	c.cache.Add(url, data)
	return locationPokemonResp, nil
}
