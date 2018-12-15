package list

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

func Swap(n, m *Node) {
	nv := n.Value
	n.Value = m.Value
	m.Value = nv
}
