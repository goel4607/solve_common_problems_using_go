package Q11_findBalancedStrWithMinBrackets

type Solution2 struct {
}

func (s2 Solution2) findBalancedStrWithMinBrackets(s string) string {
	// declare a stack of startBrackets indexes say sBktIdxs
	var sBktIdxs Stack
	var eBktIdxs []int
	// iterate over the stack
	for i, v := range s {
		if v == '(' {
			// if the element is a start bracket then push its position on sBktIdxs and append it into r
			sBktIdxs.Push(i)
		} else if v == ')' { // else if the element is end bracket
			// then if stack is non-empty then pop once, append it into r
			if !sBktIdxs.Empty() {
				sBktIdxs.Pop()
			} else {
				// else continue the loop without appending it into r
				eBktIdxs = append(eBktIdxs, i)
			}
		}
	}

	if sBktIdxs.Len() == 0 && len(eBktIdxs) == 0 {
		return s
	}

	// combine unwanted braket indexes
	unwantedBrks := make(map[int]bool)
	for !sBktIdxs.Empty() {
		unwantedBrks[sBktIdxs.Pop()] = true
	}

	for _, v := range eBktIdxs {
		unwantedBrks[v] = true
	}

	var r []rune
	for i, v := range s {
		if ok := unwantedBrks[i]; !ok {
			r = append(r, v)
		}
	}

	// make a string from slice
	return string(r)
}
