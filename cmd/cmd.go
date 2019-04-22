package cmd

import (
	"bufio"
	"strings"
)

func ReadString(reader *bufio.Reader) string {
	command, _ := reader.ReadString('\n')
	command = strings.TrimRight(command, "\n")
	return command
}
