package singlylinkedlist

type Node struct {
	Next *Node
	Data any
}

func (n *Node) Init(d []any) {
	p := n
	for i, d := range d {
		if i == 0 {
			n.Data = d // using the current node as the first node
		} else {
			t := &Node{
				Data: d,
				Next: nil,
			}
			p.Next = t
			p = t
		}
	}
}

func (n *Node) AllData() []any {
	all := make([]any, 0)

	t := n
	for t != nil {
		all = append(all, t.Data)
		t = t.Next
	}

	return all
}
