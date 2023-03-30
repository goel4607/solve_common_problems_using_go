package Q7_MNReverser

import (
	"github.com/solve_common_problems_using_go"
)

//start time: 10:48 pm
//end time: 11:37 pm

type Prac23Week11 struct {
}

func (p Prac23Week11) reverse(head *solve_common_problems_using_go.SLLNode, m, n int) *solve_common_problems_using_go.SLLNode {
	//1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7, 	m=3, n=5
	//1 -> 2 ->
	//save prev = 2
	//save next = 6
	//3 -> 6
	//4 -> 3

	//iterate till the n nodes is met, store m-1 and n pointers in say: first=m-1 and last=n
	var first, last *solve_common_problems_using_go.SLLNode

	var count int
	t := head
	for t != nil {
		count++
		if count == m-1 {
			first = t
		} else if count == n {
			last = t.Next
			break
		}
		t = t.Next
	}

	//curr = first->next
	var curr *solve_common_problems_using_go.SLLNode
	if first != nil {
		curr = first.Next
	} else {
		curr = head
	}
	//prev = last
	prev := last
	//start iterating from curr node till it is != to last
	for curr != last {
		//save curr -> next as t => t = 4 => t = 5 => 6
		t := curr.Next
		//set curr.Next = prev => 3 -> 6 => 4 -> 3 => 5 -> 4
		curr.Next = prev
		//prev = curr => 3 => 4 => 5
		prev = curr
		//curr = t => 4 => 5 => 6
		curr = t
	}

	//if first is nil then return prev
	if first == nil {
		return prev
	}

	//otherwise first->prev => 2 -> 5 and return head
	first.Next = prev
	return head
}
