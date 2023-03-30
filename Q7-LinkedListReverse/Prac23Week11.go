package Q7_LinkedListReverse

import (
	"github.com/solve_common_problems_using_go"
)

//start time: 10:00 pm
//end time: 10:22 pm

type Prac23Week11 struct {
}

func (p Prac23Week11) reverse(n *solve_common_problems_using_go.SLLNode) *solve_common_problems_using_go.SLLNode {
	//1 -> 2 -> 3 -> 4 -> 5 => 5 -> 4 -> 3 -> 2 -> 1
	//1 -> nil
	//init prev = nil
	var prev *solve_common_problems_using_go.SLLNode
	//curr = n
	curr := n
	//while curr != nil
	for curr != nil {
		//save curr -> next as temp i.e. temp = 2
		t := curr.Next
		//curr -> prev 1 -> nil
		curr.Next = prev
		//prev = curr => prev = 1
		prev = curr
		//curr = temp => 2
		curr = t
	}
	//return prev
	return prev
}
