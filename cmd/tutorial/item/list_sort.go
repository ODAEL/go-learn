package item

import (
	"bufio"
	"fmt"
	"go-learn/cmd/iohelper"
	"go-learn/cmd/list"
	"os"
	"strconv"
	"strings"
)

func StartListSort() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Create list by writing numbers using Enter and print 's' to sort the list. Use 'q' to quit")

	myList := list.List{}

	for {
		command := iohelper.ReadString(reader)
		num, err := strconv.Atoi(command)

		switch {
		case err == nil:
			myList.Append(num)
			fmt.Println(myList.GetSlice())
			break

		case strings.Compare(command, "q") == 0:
			return
			break

		case strings.Compare(command, "s") == 0:
			myList = *myList.QSort()
			fmt.Println(myList.GetSlice())
			break

		default:
			fmt.Println("Try another command")
			break
		}
	}
}
