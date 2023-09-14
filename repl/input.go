package repl

import (
	"bufio"
	"strings"
)

func cleanInput(reader *bufio.Reader) (command string, arg string, err error) {
	output, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}
	output = strings.TrimSpace(output)
	output = strings.ToLower(output)
	tokens := strings.Split(output, " ")
	var cmd string
	var args string
	cmd = tokens[0]
	if len(tokens) > 1 {
		args = tokens[1]
	}
	return cmd, args, nil
}
