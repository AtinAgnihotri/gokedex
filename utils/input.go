package utils

import (
	"bufio"
	"strings"
)

func cleanInput(reader *bufio.Reader) (string, error) {
	output, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	output = strings.TrimSpace(output)
	output = strings.ToLower(output)
	return output, nil
}
