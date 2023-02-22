package rightSideView

import "github.com/solve_common_problems_using_go/bst"

type RightSideViewCombiner interface {
	combineRightSideView(n *bst.Node) []int
}
