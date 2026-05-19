package main

import (
	"github.com/ianivr/pokedex-go/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient()
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
