package Q15_MaxDepthOfBinaryTree

import (
	"github.com/solve_common_problems_using_go"
)

type S1MaxDepthFinder struct {
}

func (s1 S1MaxDepthFinder) findMaxDepth(n *solve_common_problems_using_go.BSTNode) int {
	if n == nil {
		return 0
	}

	return s1.findMaxDepthUsingInorder(n, 1)
}

func (s1 S1MaxDepthFinder) findMaxDepthUsingInorder(n *solve_common_problems_using_go.BSTNode, max int) int {
	if n == nil {
		return max
	}

	var maxL, maxR int

	if n.Left != nil {
		maxL = s1.findMaxDepthUsingInorder(n.Left, max+1)
	}

	if n.Right != nil {
		maxR = s1.findMaxDepthUsingInorder(n.Right, max+1)
	}

	if max > maxL && max > maxR {
		return max
	} else if maxL < maxR {
		return maxR
	} else {
		return maxL
	}
}
