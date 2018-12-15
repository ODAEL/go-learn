package list

func (l *List) Sort() *List {
	ls := l.sortGetSeparatedNodes()

	return sortRecursiveMerge(ls)
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

func sortRecursiveMerge(ls []*List) *List {
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

		ll = append(ll, sortMergeSorted(f, s))
	}

	return sortRecursiveMerge(ll)
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