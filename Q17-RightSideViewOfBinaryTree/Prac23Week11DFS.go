package Q17_RightSideViewOfBinaryTree

import (
	"github.com/solve_common_problems_using_go"
)

type Prac23Week11DFS struct {
}

func (p Prac23Week11DFS) combineRightSideView(n *solve_common_problems_using_go.BSTNode) []int {
	if n == nil {
		return []int{}
	}

	return p.populateOnlyWhenLevelIsEmpty(n, 0, make([]int, 0))
}

func (p Prac23Week11DFS) populateOnlyWhenLevelIsEmpty(n *solve_common_problems_using_go.BSTNode, level int, rhsView []int) []int {
	if len(rhsView) == level {
		rhsView = append(rhsView, n.Data)
	}

	if n.Right != nil {
		rhsView = p.populateOnlyWhenLevelIsEmpty(n.Right, level+1, rhsView)
	}

	if n.Left != nil {
		rhsView = p.populateOnlyWhenLevelIsEmpty(n.Left, level+1, rhsView)
	}

	return rhsView
}
