package Q11_findBalancedStrWithMinBrackets

type Solution1 struct {
}

func (s1 Solution1) findBalancedStrWithMinBrackets(s string) string {
	// declare Stack of int with each value rep the position of the bracket
	posStack := Stack{}
	// declare another Stack for dangling back bracket positions
	backBracs := []int{}
	// store the pos of starting brackets in a Stack
	// loop over the given string one char at a time
	for i, v := range s {
		// if the char is a starting bracket then store that into the Stack
		if v == '(' {
			posStack.Push(i)
		} else if v == ')' {
			// if the char is an ending bracket then if the Stack is not empty then Pop one element of the Stack
			if posStack.Empty() {
				backBracs = append(backBracs, i)
			} else {
				posStack.Pop()
			}
		}
	}
	// make a map of all the remaining positions in both slice and Stack
	if posStack.Empty() && len(backBracs) == 0 {
		return s
	}

	unWantedBracs := make(map[int]bool)
	for !posStack.Empty() {
		unWantedBracs[posStack.Pop()] = true
	}

	for _, v := range backBracs {
		unWantedBracs[v] = true
	}
	// create a slice of chars which will hold the balanced chars
	balStr := make([]rune, 0, len(s)-len(backBracs)-posStack.Len())
	// loop over the given string from backwords
	for i, r := range s {
		// if the index matches with the map index then delete the map value
		if _, ok := unWantedBracs[i]; !ok {
			// otherwise add that char into the balanced chars slice
			balStr = append(balStr, r)
		}
	}
	// return the balanced chars slice
	return string(balStr)
}

type Stack struct {
	values []int
}

func (s *Stack) Push(a int) {
	s.values = append(s.values, a)
}

func (s *Stack) Pop() int {
	v := s.values[len(s.values)-1]
	s.values = s.values[0 : len(s.values)-1]

	return v
}

func (s *Stack) Peek() (int, bool) {
	if s.Empty() {
		return 0, false
	}
	return s.values[len(s.values)-1], true
}

func (s *Stack) Empty() bool {
	if len(s.values) == 0 {
		return true
	}

	return false
}

func (s *Stack) Len() int {
	return len(s.values)
}
