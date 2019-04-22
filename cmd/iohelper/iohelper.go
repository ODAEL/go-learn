package iohelper

import (
	"bufio"
	"strings"
)

func ReadString(reader *bufio.Reader) string {
	str, _ := reader.ReadString('\n')
	str = strings.TrimRight(str, "\n")
	return str
}
