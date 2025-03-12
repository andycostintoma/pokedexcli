# Pokedex CLI

Welcome to the Pokedex CLI, a command-line interface for exploring and catching Pokémon using the PokeAPI. This project allows users to interact with Pokémon data, catch Pokémon, and manage their collection through a simple command-line interface.

## Features

- **Catch Pokémon**: Attempt to catch Pokémon by providing their names.
- **Explore Locations**: Discover Pokémon in various locations.
- **Inspect Pokémon**: View details about the Pokémon you've caught.
- **Pokedex**: List all the Pokémon you've caught.
- **Help**: Get a list of available commands and their descriptions.
- **Exit**: Close the application.

## Requirements

- Go 1.23.2
- Internet connection (to access the PokeAPI)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/andycostintoma/pokedexcli.git
   cd pokedexcli
   ```

2. Build the application:

   ```bash
   go build -o pokedexcli
   ```

3. Run the application:

   ```bash
   ./pokedexcli
   ```

## Usage

Once the application is running, you can use the following commands:

- `help`: Displays a help message with available commands.
- `map`: Get the next page of locations.
- `mapb`: Get the previous page of locations.
- `explore <location_name>`: Explore a location to find Pokémon.
- `catch <pokemon_name>`: Attempt to catch a Pokémon.
- `inspect <pokemon_name>`: View details about a caught Pokémon.
- `pokedex`: See all the Pokémon you've caught.
- `exit`: Exit the Pokedex.

### Example Commands

- To catch a Pokémon named "pikachu":

  ```
  catch pikachu
  ```

- To explore a location named "forest":

  ```
  explore forest
  ```

- To inspect a caught Pokémon named "bulbasaur":

  ```
  inspect bulbasaur
  ```
## Acknowledgments

- [PokeAPI](https://pokeapi.co/) for providing the Pokémon data.

