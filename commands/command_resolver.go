package commands

import (
	types "github.com/AtinAgnihotri/gokedex/types"
)

func GetCommand(cmd string) types.CliCommand {
	return getCommandMap()[cmd]
}
