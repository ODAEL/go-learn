package tutorial

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"studyGo/cmd"
)

var reader *bufio.Reader

const greetingText = "Hello! Please, select the example to start:"

func Start() {
	reader = bufio.NewReader(os.Stdin)

	cmd.Println(greetingText)
	cmd.EmptyLine()

	printExamplesMap()
	printListHint()

	for {
		command := cmd.ReadString(reader);
		exampleNumber, err := strconv.Atoi(command);
		switch {
		case err == nil:
			if item, mapErr := examples[exampleNumber]; mapErr {
				item.startFunction()
			} else {
				cmd.Println("There is no such example")
			}
			break
		case strings.Compare(command, "q") == 0:
			cmd.Println("Bye!")
			return
			break
		case strings.Compare(command, "l") == 0:
			printExamplesMap()
			break
		default:
			cmd.Println("Try another command")
			break
		}
	}
}
