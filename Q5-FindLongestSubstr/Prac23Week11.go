package Q5_FindLongestSubstr

type Prac23Week11 struct {
}

func (p Prac23Week11) findLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	//keep a map of char vs its first index
	lastSeenIdxByChar := make(map[rune]int)
	//define wLeft and wRight variables
	var wLeft, wRight, max int
	//iterate over the input char by char
	for i, v := range s {
		wRight++
		//check if the char is not present in map then add it with its value as current index
		if idx, ok := lastSeenIdxByChar[v]; !ok {
			lastSeenIdxByChar[v] = i
		} else if idx < wLeft { //if present, then check if its index is < window left position,
			//if yes, then do not modify the window left position and update the char index to current index
			lastSeenIdxByChar[v] = i
		} else {
			//if no, then modify the window left to its old position + 1
			wLeft = idx + 1
			lastSeenIdxByChar[v] = i
		}

		newMax := wRight - wLeft
		if newMax > max {
			max = newMax
		}
	}

	return max
}
