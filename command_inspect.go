package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, parameters ...string) error {
	if len(parameters) == 0 {
		return errors.New("inspect requires a pokemon name")
	}

	pokemon, exists := cfg.caughtPokemon[parameters[0]]
	if !exists {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf("- %s\n", t.Type.Name)
	}

	return nil
}
