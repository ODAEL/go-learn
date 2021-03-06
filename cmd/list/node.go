package list

import "fmt"

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

func (node *Node) SetPrev(prev *Node) {
	node.Prev = prev

	if prev != nil {
		prev.Next = node
	}
}

func (node *Node) SetNext(next *Node) {
	node.Next = next

	if next != nil {
		next.Prev = node
	}
}

func SwapValues(n, m *Node) {
	nv := n.Value
	n.Value = m.Value
	m.Value = nv
}

func Swap(n, m *Node) {
	if n.Next == m {
		swapNear(n, m)
		return
	}

	if m.Next == n {
		swapNear(m, n)
		return
	}

	swapSeparate(n, m)
}

func swapSeparate(n, m *Node) {
	np := n.Prev
	nn := n.Next
	mp := m.Prev
	mn := m.Next

	n.Prev = mp
	n.Next = mn
	m.Prev = np
	m.Next = nn

	if mp != nil {
		mp.Next = n
	}

	if mn != nil {
		mn.Prev = n
	}

	if np != nil {
		np.Next = m
	}

	if nn != nil {
		nn.Prev = m
	}
}

func swapNear(n, m *Node) {
	np := n.Prev
	mn := m.Next

	n.Prev = m
	m.Next = n

	n.Next = mn
	m.Prev = np

	if np != nil {
		np.Next = m
	}

	if mn != nil {
		mn.Prev = n
	}
}

func (node Node) String() string {
	return fmt.Sprintf("%v", node.Value)
}
