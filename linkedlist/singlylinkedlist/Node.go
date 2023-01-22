package singlylinkedlist

type Node struct {
	next *Node
	data any
}

func (n *Node) Init(d []any) {
	p := n
	for _, d := range d {
		t := &Node{
			data: d,
			next: nil,
		}

		p.next = t
	}
}

func (n *Node) AllData() []any {
	all := make([]any, 0)

	t := n
	for t.next != nil {
		all = append(all, t.data)
		t = t.next
	}

	return all
}
