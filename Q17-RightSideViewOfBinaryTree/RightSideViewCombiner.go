package Q17_RightSideViewOfBinaryTree

import (
	"github.com/solve_common_problems_using_go"
)

type RightSideViewCombiner interface {
	combineRightSideView(n *solve_common_problems_using_go.BSTNode) []int
}
