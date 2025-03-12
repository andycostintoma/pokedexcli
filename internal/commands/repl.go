package commands

import (
	"bufio"
	"fmt"
	"github.com/andycostintoma/pokedexcli/internal/config"
	"os"
	"strings"
)

func StartRepl(cfg *config.Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		var args []string
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := GetCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	Name        string
	Description string
	callback    func(*config.Config, ...string) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			callback:    CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Get the next page of locations",
			callback:    CommandMapf,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Get the previous page of locations",
			callback:    CommandMapb,
		},
		"explore": {
			Name:        "explore <location_name>",
			Description: "Explore a location",
			callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch <pokemon_name>",
			Description: "Attempt to catch a pokemon",
			callback:    CommandCatch,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "See all the pokemon you've caught",
			callback:    CommandPokedex,
		},
		"inspect": {
			Name:        "inspect <pokemon_name>",
			Description: "View details about a caught Pokemon",
			callback:    CommandInspect,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			callback:    CommandExit,
		},
	}
}
