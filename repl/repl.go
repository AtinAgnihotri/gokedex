package repl

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Repl(signal chan int) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("gokedex > ")
		cmdStr, arg, err := cleanInput(reader)
		if err != nil {
			log.Fatal("Gokeded:", err)
		}
		cmd, err := GetCommand(cmdStr)
		if err != nil {
			continue
		}
		cmdErr := cmd.Callback(arg)
		if cmdErr != nil {
			log.Fatal("Gokeded", cmdErr)
		}
		if cmdStr == "exit" {
			break
		}
	}
	signal <- 0
}
