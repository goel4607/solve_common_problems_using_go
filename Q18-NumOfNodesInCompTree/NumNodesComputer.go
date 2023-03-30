package Q18_NumOfNodesInCompTree

import (
	"github.com/solve_common_problems_using_go"
)

type NumNodesComputer interface {
	ComputeNumNodes(n *solve_common_problems_using_go.BSTNode) int
}
