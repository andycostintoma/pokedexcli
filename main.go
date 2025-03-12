package main

import (
	"github.com/andycostintoma/pokedexcli/internal/commands"
	"github.com/andycostintoma/pokedexcli/internal/config"
	"time"

	"github.com/andycostintoma/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config.Config{
		CaughtPokemon: map[string]pokeapi.Pokemon{},
		PokeapiClient: pokeClient,
	}

	commands.StartRepl(cfg)
}
