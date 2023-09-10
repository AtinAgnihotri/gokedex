package commands

import (
	types "github.com/AtinAgnihotri/gokedex/types"
)

func getCommandMap() map[string]types.CliCommand {
	return map[string]types.CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback: func() error {
				return nil
			},
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback: func() error {
				return nil
			},
		},
	}
}
