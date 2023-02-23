package numNodesInCompBst

import (
	"github.com/solve_common_problems_using_go/bst"
	"math"
)

type Solution1NumNodesComputer struct {
}

func (s Solution1NumNodesComputer) computeNumNodes(n *bst.Node) int {
	// find the height of the tree. If only one node then height = 0
	h := s.computeHeight(n)
	if h == -1 {
		return 0
	} else if h == 0 {
		return 1
	}
	// max nodes at height = 2 pow h, so left = 0 and right = max - 1
	maxNodesAtLastLvl := int(math.Pow(2, float64(h)))
	ttlNodesTillLastLvl := maxNodesAtLastLvl - 1
	// total nodes at height - 1 = max - 1
	// plus the nodes at the height level
	// we can use binary search to find how many nodes are at the last level with left = 0 and right = max - 1
	left := 0
	right := maxNodesAtLastLvl - 1
	for left < right {
		// mid = left + right / 1 rounded to ceiling
		mid := int(math.Ceil(float64(left+right) / 2))
		// check if mid exists
		if s.checkNodeExists(n, h, mid) {
			// if yes then left = mid + 1
			left = mid
		} else { // else right = left - 1
			right = mid - 1
		}
	}
	// return total + left
	return ttlNodesTillLastLvl + left + 1
}

func (s Solution1NumNodesComputer) computeHeight(n *bst.Node) int {
	h := -1
	for n != nil {
		n = n.Left
		h++
	}

	return h
}

func (s Solution1NumNodesComputer) checkNodeExists(n *bst.Node, height, mid int) bool {
	// for count < height
	left, right := 0, int(math.Pow(2, float64(height))-1)
	var count int
	for count < height {
		// nowMid := ceil (right + left / 2) say 4
		m := int(math.Ceil(float64(right+left) / 2))
		if mid >= m {
			// then n = n.Right, left = mid
			n = n.Right
			left = m
		} else {
			// else n = n.Left, right = mid-1
			n = n.Left
			right = m - 1
		}

		count++
	}

	// check if n is not nil then true
	if n != nil {
		return true
	} else {
		return false
	}
}
