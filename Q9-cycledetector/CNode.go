package Q9_cycledetector

type CNode struct {
	Next *CNode
	Data any
}

func (n *CNode) Init(d []any) {
	p := n

	seenNodes := make(map[any]*CNode)

	for i, d := range d {
		if i == 0 {
			n.Data = d // using the current node as the first node
			seenNodes[d] = n
		} else if seen, ok := seenNodes[d]; !ok {
			t := &CNode{
				Data: d,
				Next: nil,
			}
			seenNodes[d] = t
			p.Next = t
			p = t
		} else {
			p.Next = seen
			break
		}
	}
}

func (n *CNode) AllData() []any {
	all := make([]any, 0)

	seenNodes := make(map[any]bool)
	t := n
	for t != nil {
		all = append(all, t.Data)
		if _, ok := seenNodes[t.Data]; !ok {
			seenNodes[t.Data] = true
			t = t.Next
		} else {
			break // cycle detected
		}
	}

	return all
}
