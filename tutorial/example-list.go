package tutorial

import (
	"go-learn/cmd"
	"go-learn/tutorial/example"
	"strconv"
)

var examples = map[int]exampleItem{
	1: exampleItem{
		"list-sort",
		"Simple example of quick sort of the list",
		example.StartList,
	},
}

func printExamplesMap() {
	for i := 1; i <= len(examples); i++ {
		cmd.Println(strconv.Itoa(i) + ". " + examples[i].description)
	}

	cmd.EmptyLine()
}
