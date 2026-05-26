package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, parameters ...string) error {
	if len(parameters) == 0 {
		return errors.New("explore requires a location name or id")
	}

	resp, err := cfg.pokeapiClient.ListLocationPokemon(parameters[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", resp.Name)
	fmt.Printf("Found Pokemon:\n")
	for _, encounter := range resp.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}
