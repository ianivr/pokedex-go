package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next     *string
	previous *string
}

type apiResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    help,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    exit,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 locations of the pokedex",
			callback:    mapForward,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 locations of the pokedex",
			callback:    mapBack,
		},
	}
}

func exit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func help(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func fetchLocations(cfg *config, url string) error {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	var maps apiResponse
	err = json.Unmarshal(body, &maps)
	if err != nil {
		return err
	}
	cfg.next = maps.Next
	cfg.previous = maps.Previous

	for _, result := range maps.Results {
		fmt.Printf("%s\n", result.Name)
	}
	return nil
}

func mapForward(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg != nil {
		if cfg.next != nil {
			url = *cfg.next
		}
	}
	return fetchLocations(cfg, url)
}

func mapBack(cfg *config) error {
	if cfg == nil || cfg.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	return fetchLocations(cfg, *cfg.previous)
}
