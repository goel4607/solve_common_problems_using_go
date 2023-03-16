package numNodesInCompBst

import (
	"github.com/solve_common_problems_using_go/bst"
	"math"
)

type Prac23Week11 struct {
}

func (p Prac23Week11) ComputeNumNodes(n *bst.Node) int {
	//calc height of tree
	h := p.findHeight(n, 1)
	//total nodes = 2^4 - 1
	//calc nodes at last level by using formula: 2^n = 2^3= 8
	maxNodesAtLastLvl := int(math.Pow(float64(2), float64(h-1)))
	//the last level nodes range: 0 to 2^(n-1) -1 (since we are starting from zero)
	start, end := 0, maxNodesAtLastLvl-1
	//implement binary search by considering the number of nodes inc by 1
	//loop till start <= end
	for start < end {
		//calc mid = (end - start) / 2 + 1
		mid := int(math.Ceil(float64(start+end) / 2))
		//find if mid exists
		if p.checkIfElementExists(n, mid, start, end, 1, h) {
			//if yes, then change start = mid
			start = mid
		} else {
			//otherwise, change end = mid - 1
			end = mid - 1
		}
	}

	maxNodesBeforeLastLvl := int(math.Pow(float64(2), float64(h-1)) - 1)
	//calc num nodes excluding the last level: 2^(n-1) - 1
	//add mid element to it
	//return the number
	return maxNodesBeforeLastLvl + start + 1
}

func (p Prac23Week11) findHeight(n *bst.Node, level int) int {
	if n.Left != nil {
		level = p.findHeight(n.Left, level+1)
	}

	return level
}

func (p Prac23Week11) checkIfElementExists(n *bst.Node, num, start, end, level, h int) bool {
	if level == h { // last level
		if n != nil {
			return true
		} else {
			return false
		}
	}

	mid := int(math.Ceil(float64(start+end) / float64(2)))
	if num < mid {
		return p.checkIfElementExists(n.Left, num, start, mid-1, level+1, h)
	} else {
		return p.checkIfElementExists(n.Right, num, mid, end, level+1, h)
	}
}
