package findBalancedStrWithMinBrackets

// start time: 12:14 am
// end time: 12:45 pm

type Prac23Week11 struct {
}

func (p Prac23Week11) findBalancedStrWithMinBrackets(s string) string {
	bracks := MyStack{}
	ignoredPos := make(map[int]bool)
	chars := make([]rune, 0)
	for i, v := range s {
		switch v {
		case '(':
			bracks.push(i)
		case ')':
			if bracks.isEmpty() {
				ignoredPos[i] = true
			} else {
				bracks.pop()
			}
		}

		chars = append(chars, v)
	}

	for !bracks.isEmpty() {
		idx := bracks.pop()
		chars[idx] = -1
	}

	bal := make([]rune, 0)
	for i, c := range chars {
		if c == -1 {
			continue
		}

		if _, ok := ignoredPos[i]; ok {
			continue
		}

		bal = append(bal, c)
	}

	return string(bal)
}

type MyStack []int

func (s *MyStack) push(v int) {
	*s = append(*s, v)
}

func (s *MyStack) pop() int {
	origLen := len(*s)
	v := (*s)[origLen-1]
	*s = (*s)[:origLen-1]
	return v
}

func (s *MyStack) isEmpty() bool {
	if len(*s) == 0 {
		return true
	}

	return false
}
