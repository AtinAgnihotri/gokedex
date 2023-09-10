package commands

import (
	"fmt"

	types "github.com/AtinAgnihotri/gokedex/types"
)

func GetCommand(cmdStr string) types.CliCommand {
	cmd, ok := getCommandMap()[cmdStr]
	if !ok {
		fmt.Println(fmt.Sprintf("Unknown Command %v. Try using >help", cmdStr))
	}
	return cmd
}
