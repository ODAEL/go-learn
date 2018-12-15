package tutorial

import (
	"strconv"
	"studyGo/cmd"
	"studyGo/tutorial/example"
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