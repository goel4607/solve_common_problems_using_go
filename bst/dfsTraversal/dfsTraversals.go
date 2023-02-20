package dfsTraversal

import "github.com/solve_common_problems_using_go/bst"

type DfsTraversals interface {
	Preorder(*bst.Node) []int
	Inorder(*bst.Node) []int
	Postorder(*bst.Node) []int
}
