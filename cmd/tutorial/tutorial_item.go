package tutorial

import (
	"fmt"
	"go-learn/cmd/tutorial/item"
	"strconv"
)

type exampleItem struct {
	name          string
	description   string
	startFunction func()
}

func (item exampleItem) start() {
	// TODO
}

var examples = map[int]exampleItem{
	1: exampleItem{
		"list-sort",
		"Simple item of quick sort of the list",
		item.StartListSort,
	},
	2: exampleItem{
		"panic-attack",
		"Panic Attack The Game",
		item.StartPanicAttack,
	},
	3: exampleItem{
		"rot13",
		"Encode string with rot13 encoding",
		item.StartRot13,
	},
}

func printExamplesMap() {
	for i := 1; i <= len(examples); i++ {
		fmt.Println(strconv.Itoa(i) + ". " + examples[i].description)
	}

	fmt.Println()
}
