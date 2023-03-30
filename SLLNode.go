package solve_common_problems_using_go

type SLLNode struct {
	Next *SLLNode
	Data any
}

func (n *SLLNode) Init(d []any) {
	p := n
	for i, d := range d {
		if i == 0 {
			n.Data = d // using the current node as the first node
		} else {
			t := &SLLNode{
				Data: d,
				Next: nil,
			}
			p.Next = t
			p = t
		}
	}
}

func (n *SLLNode) AllData() []any {
	all := make([]any, 0)

	t := n
	for t != nil {
		all = append(all, t.Data)
		t = t.Next
	}

	return all
}
