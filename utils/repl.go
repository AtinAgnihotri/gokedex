package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/AtinAgnihotri/gokedex/commands"
)

func Repl() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("gokedex > ")
		cmdStr, err := cleanInput(reader)
		if err != nil {
			log.Fatal("Gokedead:", err)
		}
		cmd := commands.GetCommand(cmdStr)
		fmt.Println(cmd.Description)
		if cmdStr == "exit" {
			break
		}
	}
}
