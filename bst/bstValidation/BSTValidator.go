package bstValidation

import "github.com/solve_common_problems_using_go/bst"

type BSTValidator interface {
	Validate(n *bst.Node) bool
}
