package solve_common_problems_using_go

type Queue struct {
	vals []*BSTNode
}

func (q *Queue) Enqueue(n *BSTNode) {
	q.vals = append(q.vals, n)
}

func (q *Queue) Dequeue() *BSTNode {
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

func (q *Queue) Peek() *BSTNode {
	if q.Empty() {
		return nil
	}

	return q.vals[0]
}

func (q *Queue) Len() int {
	return len(q.vals)
}
