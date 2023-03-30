package Q17_RightSideViewOfBinaryTree

import (
	"github.com/solve_common_problems_using_go"
)

type Prac23Week11BFS struct {
}

func (p Prac23Week11BFS) combineRightSideView(n *solve_common_problems_using_go.BSTNode) []int {
	if n == nil {
		return []int{}
	}

	rhsView := make([]int, 0)

	q := make([]*solve_common_problems_using_go.BSTNode, 0)
	q = append(q, n)

	for len(q) > 0 {
		count := len(q)

		var last int
		for count > 0 {
			e := q[0]
			q = q[1:]

			if e.Left != nil {
				q = append(q, e.Left)
			}

			if e.Right != nil {
				q = append(q, e.Right)
			}

			last = e.Data
			count--
		}

		rhsView = append(rhsView, last)
	}

	return rhsView
}
