package list

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const COMMAND_LIST  =

func Start() {
	reader := bufio.NewReader(os.Stdin)

	for {
		command, _ := reader.ReadString('\n')
		command = strings.TrimRight(command, "\n")

		switch {
		case strings.Compare(command, "q") == 0:
			cmdPrint("Bye")
			return
		}
	}
}

func cmdPrint(text string) {
	fmt.Println(text)
}

func getCommand() int {

}