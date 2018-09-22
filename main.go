package main

import (
	"fmt"
	"studyGo/mylist"
)

func main() {
	l := mylist.MakeBySlice([]int {1, 5, 4, 54, 2, 6, 3, 543, 2324, 432, 8})


	fmt.Println(l.Sort().GetSlice())
}
