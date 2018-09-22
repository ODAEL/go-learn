package mylist

// ------------- Node -------------

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

func (n *Node) SetPrev(m *Node) {
	n.Prev = m
	m.Next = n
}

func (n *Node) SetNext(m *Node) {
	n.Next = m
	m.Prev = n
}

func Swap(n, m *Node)  {
	nv := n.Value
	n.Value = m.Value
	m.Value = nv
}

// ------------- List -------------

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

func (l *List) GetSlice() []int  {
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
	for _, v := range s  {
		l.Append(v)
	}
	return &l
}

// ------------- Sort -------------

func (l *List) Sort() *List {
	ls := l.sortGetSeparatedNodes()

	return sortRecurMerge(ls)
}

func (l *List) sortGetSeparatedNodes() []*List {
	s := make([]*List, 0)
	for n := l.First; n != nil; {
		m := n
		n = n.Next

		m.Next = nil
		if n != nil {
			n.Prev = nil
		}

		s = append(s, &List{m})
	}
	return s
}

func sortRecurMerge(ls []*List) *List {
	l := len(ls)
	if l == 1 {
		return ls[0]
	}

	ll := make([]*List, 0)
	for i := 0; i < l; i += 2 {

		f := ls[i]
		var s = &List{}
		if i + 1 == l {
			s = &List{}
		} else {
			s = ls[i + 1]
		}

		ll = append(ll, sortMergeSorted(f, s))
	}

	return sortRecurMerge(ll)
}

func sortMergeSorted(l, p *List) *List {
	f := List{}
	for ; p.First != nil || l.First != nil; {

		if p.First == nil {
			f.Append(l.PopFirst())
			continue
		}

		if l.First == nil {
			f.Append(p.PopFirst())
			continue
		}

		if p.First.Value > l.First.Value {
			f.Append(l.PopFirst())
		} else {
			f.Append(p.PopFirst())
		}

	}
	return &f
}