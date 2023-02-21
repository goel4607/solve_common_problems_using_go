package maxDepth

import "github.com/solve_common_problems_using_go/bst"

type MaxDepthFinder interface {
	findMaxDepth(n *bst.Node) int
}
