package bst

type Queue struct {
	vals []*Node
}

func (q *Queue) Enqueue(n *Node) {
	q.vals = append(q.vals, n)
}

func (q *Queue) Dequeue() *Node {
	v := q.vals[0]
	q.vals = q.vals[1:]
	return v
}

func (q *Queue) Empty() bool {
	if len(q.vals) == 0 {
		return true
	} else {
		return false
	}
}

func (q *Queue) Peek() *Node {
	if q.Empty() {
		return nil
	}

	return q.vals[0]
}

func (q *Queue) Len() int {
	return len(q.vals)
}
