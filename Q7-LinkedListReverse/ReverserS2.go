package Q7_LinkedListReverse

import (
	"github.com/solve_common_problems_using_go"
)

type ReverserS2Repeat struct {
}

func (s2 *ReverserS2Repeat) reverse(n *solve_common_problems_using_go.SLLNode) *solve_common_problems_using_go.SLLNode {
	// 1 -> 2 -> 3 -> 4 -> 5
	// 1 -> nil and 2 -> 1
	// prev = nil
	// curr = n
	var prev *solve_common_problems_using_go.SLLNode

	curr := n
	// while n != nil
	for curr != nil {
		// save next ptr as temp
		t := curr.Next
		// 1 -> prev
		curr.Next = prev
		// prev = curr
		prev = curr
		// curr = temp i.e 2
		curr = t
	}
	// return prev
	return prev
}
