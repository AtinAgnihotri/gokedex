package commands

import (
	"fmt"

	types "github.com/AtinAgnihotri/gokedex/types"
)

func getCommandMap() map[string]types.CliCommand {
	return map[string]types.CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback: func() error {
				fmt.Println(`Welcome to the Pokedex!
				Usage:
				
				help: Displays a help message
				exit: Exit the Pokedex`)
				return nil
			},
		},
		"exit": {
			Name:        "exit",
			Description: "Bye!",
			Callback: func() error {
				fmt.Println(``)
				return nil
			},
		},
	}
}
