package list

type List struct {
	First *Node
}

func (l *List) Append(v int) {
	n := Node{v, nil, nil}

	if l.First == nil {
		l.First = &n
		return
	}

	m := l.First
	for ; m.Next != nil; {
		m = m.Next
	}

	m.SetNext(&n)
}

func (l *List) Prepend(v int) {
	n := Node{v, nil, nil}

	if l.First == nil {
		l.First = &n
		return
	}

	m := l.First
	m.SetPrev(&n)
	l.First = &n
}

func (l *List) GetSlice() []int {
	s := make([]int, 0)

	for m := l.First; m != nil; m = m.Next {
		s = append(s, m.Value)
	}

	return s
}

func (l *List) PopFirst() int {
	if l.First == nil {
		panic("List is empty")
	}

	p := l.First
	l.First = p.Next
	if l.First != nil {
		l.First.Prev = nil
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