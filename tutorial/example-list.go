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
		example.StartListSort,
	},
	2: exampleItem{
		"panic-attack",
		"Panic Attack The Game",
		example.StartPanicAttack,
	},
}

func printExamplesMap() {
	for i := 1; i <= len(examples); i++ {
		cmd.Println(strconv.Itoa(i) + ". " + examples[i].description)
	}

	cmd.EmptyLine()
}
