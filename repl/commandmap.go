package repl

import (
	"fmt"

	"github.com/AtinAgnihotri/gokedex/api"
	"github.com/AtinAgnihotri/gokedex/helpers"
	"github.com/AtinAgnihotri/gokedex/stores"
	types "github.com/AtinAgnihotri/gokedex/types"
)

func getCommandMap() map[string]types.CliCommand {
	return map[string]types.CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback: func(cmdStr string) error {
				fmt.Println()
				defer fmt.Println()
				if len(cmdStr) == 0 {
					fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: Display map locations (In Pages)
mapb: Go back a page in map locations
explore: See all pokemons encountered in an area. Usage: explore <location>
catch: Try to catch a pokemon. Usage: catch <pokemon>
inspect: Inspect a pokemon in your pokedex. Usage: inspect <pokemon>`)
					return nil
				}
				cmd, err := GetCommand(cmdStr)
				if err != nil {
					fmt.Println(fmt.Sprintf("Unknown command %v. Try > help", cmd))
				}
				fmt.Println(fmt.Sprintf("%v: %v", cmd.Name, cmd.Description))

				return nil
			},
		},
		"exit": {
			Name:        "exit",
			Description: "Exit Gokedex",
			Callback: func(_ string) error {
				fmt.Println(``)
				return nil
			},
		},
		"map": {
			Name:        "map",
			Description: "Show a page of map locations (20 at a time)",
			Callback: func(_ string) error {
				api.GetPokeApiLocations(true)
				return nil
			},
		},
		"mapb": {
			Name:        "mapb",
			Description: "Go back a list of map locations",
			Callback: func(_ string) error {
				api.GetPokeApiLocations(false)
				return nil
			},
		},
		"explore": {
			Name:        "explore",
			Description: "See all pokemons encountered in an area\nUsage: explore <location>",
			Callback: func(location string) error {
				if len(location) == 0 {
					fmt.Printf("\nPlease give a location name as Argument\nUsage: explore <location>\n")
					return nil
				}
				api.GetPokemonsInLocation(location)
				return nil
			},
		},
		"catch": {
			Name:        "catch",
			Description: "Try to catch a pokemon\nUsage: catch <pokemon>",
			Callback: func(pokemonName string) error {
				fmt.Println()
				defer fmt.Println()
				if len(pokemonName) == 0 {
					fmt.Printf("\nPlease give a pokemon name as Argument\nUsage: catch <pokemon>\n")
					return nil
				}
				if stores.Pokedex == nil {
					stores.Pokedex = make(map[string]types.Pokemon)
				}
				_, ok := stores.Pokedex[pokemonName]
				if ok {
					fmt.Println(fmt.Sprintf("You've already caught %v!", pokemonName))
					return nil
				}
				pokemon, err := api.GetPokemon(pokemonName)
				if err != nil {
					fmt.Println(fmt.Sprintf("Couldn't find %v pokemon", pokemonName))
					return nil
				}
				fmt.Println(fmt.Sprintf("Throwing a pokeball at %v ... ", pokemonName))
				chance := helpers.GetRandom(pokemon.BaseExperience)
				if chance >= (pokemon.BaseExperience / 2) {
					fmt.Println(fmt.Sprintf("Caught %v!", pokemonName))
					if stores.Pokedex == nil {
						stores.Pokedex = make(map[string]types.Pokemon)
					}
					stores.Pokedex[pokemonName] = pokemon
				} else {
					fmt.Println(fmt.Sprintf("%v escaped!", pokemonName))
				}
				return nil
			},
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect a pokemon in your pokedex\nUsage: inspect <pokemon>",
			Callback: func(pokemonName string) error {
				fmt.Println()
				defer fmt.Println()
				if len(pokemonName) == 0 {
					fmt.Printf("\nPlease give a pokemon name as Argument\nUsage: inspect <pokemon>\n")
					return nil
				}
				if stores.Pokedex == nil {
					fmt.Println("No pokemon in your pokedex")
					return nil
				}
				pokemon, ok := stores.Pokedex[pokemonName]
				if !ok {
					fmt.Println(fmt.Sprintf("%v not in your pokedex", pokemonName))
					return nil
				}
				statsStr, typesStr := "", ""
				for idx, pokeType := range pokemon.Types {
					typesStr += "  - " + pokeType.Type.Name
					if idx < len(pokemon.Types)-1 {
						typesStr += "\n"
					}
				}
				for idx, stat := range pokemon.Stats {
					statsStr += fmt.Sprintf("  -%v: %v", stat.Stat.Name, stat.BaseStat)
					if idx < len(pokemon.Stats)-1 {
						statsStr += "\n"
					}
				}

				fmt.Println(fmt.Sprintf(`Name: %v
Height: %v
Weight: %v
Stats:
%v
Types:
%v`, pokemon.Name, pokemon.Height, pokemon.Weight, statsStr, typesStr))
				return nil
			},
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Shows all pokemon in your pokedex",
			Callback: func(_ string) error {
				fmt.Println()
				defer fmt.Println()
				emptyPokedex := "Your pokedex is empty"

				if stores.Pokedex == nil {
					fmt.Println(emptyPokedex)
					return nil
				}
				if len(stores.Pokedex) == 0 {
					fmt.Println(emptyPokedex)
					return nil
				}
				fmt.Println("Your Pokedex:")
				for pokemonName := range stores.Pokedex {
					fmt.Println(" - ", pokemonName)
				}
				return nil
			},
		},
	}
}
