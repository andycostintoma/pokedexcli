package commands

import (
	"fmt"
	"github.com/andycostintoma/pokedexcli/internal/config"
	"os"
)

func CommandExit(cfg *config.Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
