package Q12_ImplQueueWithStacks

import (
	"github.com/solve_common_problems_using_go/Q11-findBalancedStrWithMinBrackets"
)

type Solution2QUsingStack struct {
	stack1 Q11_findBalancedStrWithMinBrackets.Stack
	stack2 Q11_findBalancedStrWithMinBrackets.Stack
}

func (s *Solution2QUsingStack) enqueue(v int) {
	s.stack1.Push(v)
}

func (s *Solution2QUsingStack) dequeue() (int, bool) {
	if s.stack2.Empty() {
		if s.stack1.Empty() {
			return 0, false
		} else {
			for !s.stack1.Empty() {
				s.stack2.Push(s.stack1.Pop())
			}

			return s.stack2.Pop(), true
		}

	} else {
		return s.stack2.Pop(), true
	}
}

func (s *Solution2QUsingStack) peek() (int, bool) {
	if !s.stack2.Empty() {
		return s.stack2.Peek()
	} else if !s.stack1.Empty() {
		for !s.stack1.Empty() {
			s.stack2.Push(s.stack1.Pop())
		}

		return s.stack2.Peek()
	} else {
		return 0, false
	}
}

func (s *Solution2QUsingStack) empty() bool {
	if s.stack1.Empty() && s.stack2.Empty() {
		return true
	}

	return false
}
