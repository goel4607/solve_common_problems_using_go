package Q11_findBalancedStrWithMinBrackets

import "strings"

type Prac230601 struct {
}

func (p Prac230601) findBalancedStrWithMinBrackets(s string) string {
	//create a place holder slice for all chars except closing brackets
	chars := make([]rune, 0)
	//create a stack for matching brackets
	stack := MyStack1{}
	//iterate over all the chars in string s
	for i, v := range s {
		//check if char is '(' then push its index into stack
		if v == '(' {
			stack.Push(i)
		} else if v == ')' {
			//check if the char is ')' the check if stack is non-empty,
			//if non-empty, then pop from stack
			if stack.IsEmpty() {
				v = ' ' //replace with empty char; effectively removing the closing bracket
			} else {
				stack.Pop()
			}
		}

		chars = append(chars, v)
	}
	//end loop

	for !stack.IsEmpty() {
		i := stack.Pop()
		chars[i] = ' '
	}

	s = string(chars)
	return strings.ReplaceAll(s, " ", "")
}

type MyStack1 []int

func (ms *MyStack1) Push(v int) {
	*ms = append(*ms, v)
}

func (ms *MyStack1) Pop() int {
	oldLen := len(*ms)
	v := (*ms)[oldLen-1]
	*ms = (*ms)[0 : oldLen-1]
	return v
}

func (ms *MyStack1) Len() int {
	return len(*ms)
}

func (ms *MyStack1) IsEmpty() bool {
	if ms.Len() == 0 {
		return true
	}

	return false
}
