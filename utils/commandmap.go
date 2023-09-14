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
explore: See all pokemons encountered in an area. Usage: explore <location>`)
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
			Callback: func(arg string) error {
				GetPokemonsInLocation(arg)
				return nil
			},
		},
	}
}
