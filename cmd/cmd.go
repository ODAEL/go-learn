package cmd

import (
	"bufio"
	"fmt"
	"strings"
)

func Println(text string) {
	fmt.Println(text)
}

func Print(text string) {
	fmt.Print(text)
}

func EmptyLine() {
	fmt.Println()
}

func ReadString(reader *bufio.Reader) string {
	command, _ := reader.ReadString('\n')
	command = strings.TrimRight(command, "\n")
	return command
}
