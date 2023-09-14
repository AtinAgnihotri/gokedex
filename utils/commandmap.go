package utils

import (
	"fmt"

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
catch: Try to catch a pokemon. Usage: catch <pokemon>`)
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
				GetPokeApiLocations(true)
				return nil
			},
		},
		"mapb": {
			Name:        "mapb",
			Description: "Go back a list of map locations",
			Callback: func(_ string) error {
				GetPokeApiLocations(false)
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
				GetPokemonsInLocation(location)
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
					fmt.Printf("\nPlease give a location name as Argument\nUsage: explore <location>\n")
					return nil
				}
				pokemon, err := GetPokemon(pokemonName)
				if err != nil {
					fmt.Println(fmt.Sprintf("Couldn't find %v pokemon", pokemonName))
					return nil
				}
				fmt.Println(fmt.Sprintf("Throwing a pokeball at %v ... ", pokemonName))
				chance := GetRandom(pokemon.BaseExperience)
				if chance >= (pokemon.BaseExperience / 2) {
					fmt.Println(fmt.Sprintf("Caught %v!", pokemonName))
					if Pokedex == nil {
						Pokedex = make(map[string]types.Pokemon)
					}
					Pokedex[pokemonName] = pokemon
				} else {
					fmt.Println(fmt.Sprintf("%v escaped!", pokemonName))
				}
				return nil
			},
		},
	}
}
