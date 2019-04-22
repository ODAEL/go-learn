package tutorial

import (
	"bufio"
	"fmt"
	"go-learn/cmd"
	"os"
	"strconv"
	"strings"
)

const greetingText = "Hello! Please, select the example to start:"

const defises string = "---------------------------"

func Start() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(greetingText)
	fmt.Println()

	printExamplesMap()
	printListHint()

	for {
		command := cmd.ReadString(reader);
		exampleNumber, err := strconv.Atoi(command);
		switch {
		case err == nil:
			if item, mapErr := examples[exampleNumber]; mapErr {
				beforeStart(item.name)
				item.startFunction()
				afterEnd(item.name)
			} else {
				fmt.Println("There is no such example")
			}
			break

		case strings.Compare(command, "q") == 0:
			fmt.Println("Bye!")
			return
			break

		case strings.Compare(command, "l") == 0:
			printExamplesMap()
			break

		default:
			fmt.Println("Try another command")
			break
		}
	}
}

func printListHint() {
	fmt.Println("Print the number and press Enter to select. You can always print this list with command 'l' and quit with 'q'")
}

func beforeStart(exampleName string) {
	fmt.Println(defises + " Start: " + exampleName + " " + defises)
}

func afterEnd(exampleName string) {
	fmt.Println(defises + " End: " + exampleName + " " + defises)
}
