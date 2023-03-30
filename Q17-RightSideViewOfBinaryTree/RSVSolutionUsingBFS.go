package Q17_RightSideViewOfBinaryTree

import (
	"github.com/solve_common_problems_using_go"
)

type RSVSolutionUsingBFS struct {
}

func (s1 RSVSolutionUsingBFS) combineRightSideView(n *solve_common_problems_using_go.BSTNode) []int {

	q := solve_common_problems_using_go.Queue{}
	q.Enqueue(n)

	var values []int

	for !q.Empty() {

		var count int
		maxCount := q.Len()

		var lastVal int // last value is always the right hand node of the tree
		for count < maxCount {
			t := q.Dequeue()
			lastVal = t.Data

			if t.Left != nil {
				q.Enqueue(t.Left)
			}

			if t.Right != nil {
				q.Enqueue(t.Right)
			}

			count++
		}

		values = append(values, lastVal)
	}

	return values
}
