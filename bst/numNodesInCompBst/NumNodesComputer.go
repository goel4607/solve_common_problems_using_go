package numNodesInCompBst

import "github.com/solve_common_problems_using_go/bst"

type NumNodesComputer interface {
	ComputeNumNodes(n *bst.Node) int
}
