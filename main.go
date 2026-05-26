package main

import (
	"time"

	"github.com/ianivr/pokedex-go/internal/pokeapi"
)

type Pokemon = pokeapi.Pokemon

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: make(map[string]Pokemon),
	}
	startRepl(cfg)
}
