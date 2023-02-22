package rightSideView

import (
	"github.com/solve_common_problems_using_go/bst"
)

type RSVSolutionUsingDFSUsingArrLen struct {
}

func (s RSVSolutionUsingDFSUsingArrLen) combineRightSideView(n *bst.Node) []int {
	if n == nil {
		return []int{}
	}

	return s.visitPreorderWithRFirst(n, 0, []int{})
}

func (s RSVSolutionUsingDFSUsingArrLen) visitPreorderWithRFirst(n *bst.Node, level int, values []int) []int {
	if n == nil {
		return values
	}

	if level == len(values) {
		values = append(values, n.Data)
	}

	if n.Right != nil {
		values = s.visitPreorderWithRFirst(n.Right, level+1, values)
	}

	if n.Left != nil {
		values = s.visitPreorderWithRFirst(n.Left, level+1, values)
	}

	return values
}
