package list

import "reflect"

func (l *List) QSort() *List {
	ls := l.getNodesAsLists()

	return qsortMerge(ls)
}

func (l *List) getNodesAsLists() []*List {
	s := make([]*List, 0)
	for n := l.first; !reflect.ValueOf(n).IsNil(); {
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
