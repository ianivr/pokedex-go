package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationListResponse, error) {
	url := baseURL + "/location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	if data, ok := c.cache.Get(url); ok {
		var locationsResp LocationListResponse
		err := json.Unmarshal(data, &locationsResp)
		if err != nil {
			return LocationListResponse{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationListResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationListResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationListResponse{}, err
	}

	if res.StatusCode > 299 {
		return LocationListResponse{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
	}

	var locationsResp LocationListResponse
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return LocationListResponse{}, err
	}

	c.cache.Add(url, data)
	return locationsResp, nil
}
