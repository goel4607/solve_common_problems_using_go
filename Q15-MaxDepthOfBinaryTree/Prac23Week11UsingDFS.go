package Q15_MaxDepthOfBinaryTree

import (
	"github.com/solve_common_problems_using_go"
)

type Prac23Week11UsingDFS struct {
}

func (p Prac23Week11UsingDFS) findMaxDepth(n *solve_common_problems_using_go.BSTNode) int {
	if n == nil {
		return 0
	}

	return p.findDepth(n, 1)
}

func (p Prac23Week11UsingDFS) findDepth(n *solve_common_problems_using_go.BSTNode, d int) int {
	max := d
	if n.Left != nil {
		maxL := p.findDepth(n.Left, d+1)
		if max < maxL {
			max = maxL
		}
	}

	if n.Right != nil {
		maxR := p.findDepth(n.Right, d+1)
		if max < maxR {
			max = maxR
		}
	}

	return max
}
