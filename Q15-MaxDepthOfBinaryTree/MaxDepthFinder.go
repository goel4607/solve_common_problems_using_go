package Q15_MaxDepthOfBinaryTree

import (
	"github.com/solve_common_problems_using_go"
)

type MaxDepthFinder interface {
	findMaxDepth(n *solve_common_problems_using_go.BSTNode) int
}
