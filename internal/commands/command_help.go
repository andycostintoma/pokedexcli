package commands

import (
	"fmt"
	"github.com/andycostintoma/pokedexcli/internal/config"
)

func CommandHelp(cfg *config.Config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}
