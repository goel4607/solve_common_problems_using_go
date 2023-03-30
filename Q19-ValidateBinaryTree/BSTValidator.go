package Q19_ValidateBinaryTree

import (
	"github.com/solve_common_problems_using_go"
)

type BSTValidator interface {
	Validate(n *solve_common_problems_using_go.BSTNode) bool
}
