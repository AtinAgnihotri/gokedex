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
			Callback: func(_ string) error {
				fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: Display map locations (In Pages)
mapb: Go back a page in map locations`)
				return nil
			},
		},
		"exit": {
			Name:        "exit",
			Description: "Bye!",
			Callback: func(_ string) error {
				fmt.Println(``)
				return nil
			},
		},
		"map": {
			Name:        "map",
			Description: "Calls func",
			Callback: func(_ string) error {
				GetPokeApiLocations(true)
				return nil
			},
		},
		"mapb": {
			Name:        "mapb",
			Description: "Calls func b",
			Callback: func(_ string) error {
				GetPokeApiLocations(false)
				return nil
			},
		},
		"explore": {
			Name: "explore",
			Description: "explore pokemon in area",
			Callback: func(arg string) error {
				
			},
		}
	}
}
