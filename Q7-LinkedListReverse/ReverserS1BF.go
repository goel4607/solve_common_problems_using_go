package Q7_LinkedListReverse

import (
	"github.com/solve_common_problems_using_go"
)

type ReverserS1BF solve_common_problems_using_go.SLLNode

func (s1 *ReverserS1BF) reverse(n *solve_common_problems_using_go.SLLNode) *solve_common_problems_using_go.SLLNode {
	// 1 > 2 > 3 > 4 > 5 > nil
	// 5 > 4 > 3 > 2 > 1 > nil

	// 1 > 2 => 1 -> nil and 2 > 1
	var prev *solve_common_problems_using_go.SLLNode
	c := n
	for c != nil { // 1
		t := c.Next   // t = 2
		c.Next = prev // 1 -> nil
		prev = c      // prev = 2
		c = t
	}

	return prev
}
