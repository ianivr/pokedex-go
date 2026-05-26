package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, parameters ...string) error {
	if len(parameters) == 0 {
		return errors.New("catch requires a pokemon name or id")
	}

	resp, err := cfg.pokeapiClient.CatchPokemon(parameters[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", resp.Name)
	captureChance := rand.Intn(300)

	if captureChance > resp.BaseExperience {
		fmt.Printf("%s was caught!\n", resp.Name)
		cfg.caughtPokemon[resp.Name] = resp
		fmt.Printf("You may now inspect it with the inspect command.\n")
	} else {
		fmt.Printf("%s escaped!\n", resp.Name)
	}

	return nil
}
