package list

import (
	"fmt"
)

type List struct {
	first *Node
}

func Create(values ...int) *List {
	var next *Node
	for i := len(values) - 1; i >= 0; i-- {
		node := &Node{values[i], nil, nil}
		node.SetNext(next)
		next = node
	}

	return &List{next}
}

func (list *List) Append(v int) {
	n := Node{v, nil, nil}

	if list.first == nil {
		list.first = &n
		return
	}

	m := list.first
	for ; m.Next != nil; {
		m = m.Next
	}

	m.SetNext(&n)
}

func (list *List) Prepend(v int) {
	n := Node{v, nil, nil}

	if list.first == nil {
		list.first = &n
		return
	}

	m := list.first
	m.SetPrev(&n)
	list.first = &n
}

func (list *List) GetSlice() []int {
	var s []int

	for n := list.first; n != nil; n = n.Next {
		s = append(s, n.Value)
	}

	return s
}

func (list *List) PopFirst() int {
	if list.first == nil {
		panic("List is empty")
	}

	p := list.first
	list.first = p.Next
	if list.first != nil {
		list.first.Prev = nil
	}
	p.Next = nil

	return p.Value
}

func CreateBySlice(s []int) *List {
	return Create(s...)
}

func (list List) String() string {
	return fmt.Sprint(list.GetSlice())
}