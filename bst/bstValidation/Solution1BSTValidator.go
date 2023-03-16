package bstValidation

import (
	"github.com/solve_common_problems_using_go/bst"
	"math"
)

type Solution1BSTValidator struct {
}

func (s Solution1BSTValidator) Validate(n *bst.Node) bool {
	// using the min infinity for the left-most nodes as minimum int value
	// and the max infinity for the right-most nodes as the maximum int value
	return s.validateWithBoundaries(n, math.MinInt32, math.MaxInt32)
}

func (s Solution1BSTValidator) validateWithBoundaries(n *bst.Node, minLeft, maxRight int) bool {
	if n == nil {
		return true
	}

	if n.Left != nil {
		// left < than its parent and > than the min left
		if n.Left.Data < n.Data && minLeft < n.Left.Data {
			// for left child nodes the max limit is its parent node
			if !s.validateWithBoundaries(n.Left, minLeft, n.Data) {
				return false
			}
		} else {
			return false
		}
	}

	if n.Right != nil {
		// right > than its parent and < than the max right
		if n.Data > n.Right.Data && n.Right.Data < maxRight {
			// for right nodes min is defined by its parent
			if !s.validateWithBoundaries(n.Right, n.Data, maxRight) {
				return false
			}
		}
	}

	return true
}
