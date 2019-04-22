package list

func (l *List) BSort() *List {
	ls := l.GetSlice()

	s := MakeBySlice(ls)

	for ; bsortIteration(s) != false; {
	}

	return s
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
