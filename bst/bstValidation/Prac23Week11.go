package bstValidation

import (
	"github.com/solve_common_problems_using_go/bst"
	"math"
)

type Prac23Week11 struct {
}

func (p Prac23Week11) Validate(n *bst.Node) bool {
	return p.checkValidWithMinMax(n, math.MinInt, math.MaxInt32)
}

func (p Prac23Week11) checkValidWithMinMax(n *bst.Node, min, max int) bool {
	if n == nil {
		return true
	}

	if n.Data < min || n.Data > max {
		return false
	}

	if n.Left != nil {
		if !p.checkValidWithMinMax(n.Left, min, n.Data) {
			return false
		}
	}

	if n.Right != nil {
		if !p.checkValidWithMinMax(n.Right, n.Data, max) {
			return false
		}
	}

	return true
}
