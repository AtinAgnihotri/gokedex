package utils

import (
	"errors"
	"fmt"

	types "github.com/AtinAgnihotri/gokedex/types"
)

func GetCommand(cmdStr string) (types.CliCommand, error) {
	cmd, ok := getCommandMap()[cmdStr]
	if !ok {
		fmt.Println(fmt.Sprintf("Unknown Command %v. Try using >help", cmdStr))
		return types.CliCommand{}, errors.New("Unknown command")
	}
	return cmd, nil
}
