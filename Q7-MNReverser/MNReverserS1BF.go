package Q7_MNReverser

import (
	"github.com/solve_common_problems_using_go"
)

type MNReverserS1BF struct {
}

func (s1 *MNReverserS1BF) reverse(head *solve_common_problems_using_go.SLLNode, m, n int) *solve_common_problems_using_go.SLLNode {
	// 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7,  and m = 3 and n = 5
	// 1 -> 2 -> 5 -> 4 -> 3 -> 6 -> 7

	// iterate till you read mth element
	pos := 0
	// store the element pointing to mth element as first pos.e. first = 2
	// find the last positions next pointer and store it in last pos.e. last = 6
	var first, last *solve_common_problems_using_go.SLLNode
	t := head
	for t != nil {
		pos++

		if pos == (m - 1) {
			first = t
		} else if pos == n {
			last = t.Next
			break
		}
		t = t.Next
	}

	prev := last
	var c *solve_common_problems_using_go.SLLNode
	if first != nil {
		c = first.Next
	} else {
		c = head
	}
	// Q7-LinkedListReverse the list starting from mth element till it reaches m or nil
	pos = 0
	diff := n - m
	for c != nil && pos <= diff {
		pos++

		t = c.Next
		c.Next = prev
		prev = c
		c = t
	}

	if first != nil {
		first.Next = prev
		return head
	} else {
		return prev
	}
}
