package Q15_DFSTraversal

import (
	"github.com/solve_common_problems_using_go"
)

type DfsTraversals interface {
	Preorder(*solve_common_problems_using_go.BSTNode) []int
	Inorder(*solve_common_problems_using_go.BSTNode) []int
	Postorder(*solve_common_problems_using_go.BSTNode) []int
}
