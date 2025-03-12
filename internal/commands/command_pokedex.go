package commands

import (
	"fmt"
	"github.com/andycostintoma/pokedexcli/internal/config"
)

func CommandPokedex(cfg *config.Config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, p := range cfg.CaughtPokemon {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}
