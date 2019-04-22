package list

func (l *List) QSort() *List {
	ls := l.qsortGetNodesAsLists()

	return qsortMerge(ls)
}

func (l *List) BSort() *List {
	ls := l.GetSlice()

	s := CreateBySlice(ls)

	for ; bsortIteration(s) != false; {
	}

	return s
}

func (l *List) qsortGetNodesAsLists() []*List {
	s := make([]*List, 0)
	for n := l.first; n != nil; {
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

func qsortMerge(ls []*List) *List {
	l := len(ls)
	if l == 1 {
		return ls[0]
	}

	ll := make([]*List, 0)
	for i := 0; i < l; i += 2 {

		f := ls[i]
		var s = &List{}
		if i+1 == l {
			s = &List{}
		} else {
			s = ls[i+1]
		}

		ll = append(ll, qsortMergeSorted(f, s))
	}

	return qsortMerge(ll)
}

func qsortMergeSorted(l, p *List) *List {
	f := List{}
	for ; p.first != nil || l.first != nil; {

		if p.first == nil {
			f.Append(l.PopFirst())
			continue
		}

		if l.first == nil {
			f.Append(p.PopFirst())
			continue
		}

		if p.first.Value > l.first.Value {
			f.Append(l.PopFirst())
		} else {
			f.Append(p.PopFirst())
		}

	}
	return &f
}

func bsortIteration(l *List) (wasSorted bool) {
	wasSorted = false
	n := l.first
	for ; n.Next != nil; {
		if n.Value > n.Next.Value {
			swapNear(n, n.Next)
			wasSorted = true
		} else {
			n = n.Next
		}
	}

	for ; n.Prev != nil; {
		n = n.Prev
	}

	l.first = n

	return
}
