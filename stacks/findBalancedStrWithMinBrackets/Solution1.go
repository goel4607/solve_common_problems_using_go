package findBalancedStrWithMinBrackets

type Solution1 struct {
}

func (s1 Solution1) findBalancedStrWithMinBrackets(s string) string {
	// declare stack of int with each value rep the position of the bracket
	posStack := stack{}
	// declare another stack for dangling back bracket positions
	backBracs := []int{}
	// store the pos of starting brackets in a stack
	// loop over the given string one char at a time
	for i, v := range s {
		// if the char is a starting bracket then store that into the stack
		if v == '(' {
			posStack.push(i)
		} else if v == ')' {
			// if the char is an ending bracket then if the stack is not empty then pop one element of the stack
			if posStack.isEmpty() {
				backBracs = append(backBracs, i)
			} else {
				posStack.pop()
			}
		}
	}
	// make a map of all the remaining positions in both slice and stack
	if posStack.isEmpty() && len(backBracs) == 0 {
		return s
	}

	unWantedBracs := make(map[int]bool)
	for !posStack.isEmpty() {
		unWantedBracs[posStack.pop()] = true
	}

	for _, v := range backBracs {
		unWantedBracs[v] = true
	}
	// create a slice of chars which will hold the balanced chars
	balStr := make([]rune, 0, len(s)-len(backBracs)-posStack.len())
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

type stack struct {
	values []int
}

func (s *stack) push(a int) {
	s.values = append(s.values, a)
}

func (s *stack) pop() int {
	v := s.values[len(s.values)-1]
	s.values = s.values[0 : len(s.values)-1]

	return v
}

func (s *stack) isEmpty() bool {
	if len(s.values) == 0 {
		return true
	}

	return false
}

func (s *stack) len() int {
	return len(s.values)
}
