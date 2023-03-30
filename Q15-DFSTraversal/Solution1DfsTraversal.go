package Q15_DFSTraversal

import (
	"github.com/solve_common_problems_using_go"
)

type S1DfsTraversals struct {
	solve_common_problems_using_go.BSTNode
}

func (s1 *S1DfsTraversals) Preorder(n *solve_common_problems_using_go.BSTNode) []int {
	all := make([]int, 0)
	return s1.TraversePreorder(n, all)
}

func (s1 *S1DfsTraversals) Inorder(n *solve_common_problems_using_go.BSTNode) []int {
	all := make([]int, 0)
	return s1.TraverseInorder(n, all)
}

func (s1 *S1DfsTraversals) Postorder(n *solve_common_problems_using_go.BSTNode) []int {
	all := make([]int, 0)
	return s1.TraversePostorder(n, all)
}

func (s1 *S1DfsTraversals) TraversePreorder(n *solve_common_problems_using_go.BSTNode, all []int) []int {
	if n != nil {
		all = append(all, n.Visit())
		if n.Left != nil {
			all = s1.TraversePreorder(n.Left, all)
		}

		if n.Right != nil {
			all = s1.TraversePreorder(n.Right, all)
		}
	}

	return all
}

func (s1 *S1DfsTraversals) TraverseInorder(n *solve_common_problems_using_go.BSTNode, all []int) []int {
	if n == nil {
		return all
	}

	if n.Left != nil {
		all = s1.TraverseInorder(n.Left, all)
	}
	all = append(all, n.Data)
	if n.Right != nil {
		all = s1.TraverseInorder(n.Right, all)
	}

	return all
}

func (s1 *S1DfsTraversals) TraversePostorder(n *solve_common_problems_using_go.BSTNode, all []int) []int {
	if n == nil {
		return all
	}

	if n.Left != nil {
		all = s1.TraversePostorder(n.Left, all)
	}
	if n.Right != nil {
		all = s1.TraversePostorder(n.Right, all)
	}
	all = append(all, n.Data)

	return all
}
