package list

import "reflect"

type List struct {
	first *Node
}

func (l *List) Append(v int) {
	n := Node{v, nil, nil}

	if l.first == nil {
		l.first = &n
		return
	}

	m := l.first
	for ; m.Next != nil; {
		m = m.Next
	}

	m.SetNext(&n)
}

func (l *List) Prepend(v int) {
	n := Node{v, nil, nil}

	if l.first == nil {
		l.first = &n
		return
	}

	m := l.first
	m.SetPrev(&n)
	l.first = &n
}

func (l *List) GetSlice() []int {
	var s []int

	for n := l.first; !reflect.ValueOf(n).IsNil(); n = n.Next {
		s = append(s, n.Value)
	}

	return s
}

func (l *List) PopFirst() int {
	if l.first == nil {
		panic("List is empty")
	}

	p := l.first
	l.first = p.Next
	if l.first != nil {
		l.first.Prev = nil
	}
	p.Next = nil

	return p.Value
}

func MakeBySlice(s []int) *List {
	l := List{}
	for _, v := range s {
		l.Append(v)
	}
	return &l
}
