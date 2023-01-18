package findLongestSubstr

type LongestSubstrFinderS4Prac1 struct {
}

func (s4 LongestSubstrFinderS4Prac1) findLongestSubstring(s string) int {
	charVsIdx := make(map[rune]int, 0)
	var max, cMax, left int

	// iterate over the str from index 0 to len
	for right, v := range s {
		// if char is not present in map then inc cMax and set map[char] = index
		if c, ok := charVsIdx[v]; !ok {
			cMax++
			charVsIdx[v] = right
		} else if c < left {
			// else if the index of the char is less than left then inc max and update the map with new index
			cMax++
			charVsIdx[v] = right
		} else {
			// else if index > left then reset left = map index and cMax = current index - map index
			left = c + 1
			cMax = right - left + 1
			charVsIdx[v] = right
		}
		// if max < cMax then max = cMax
		if max < cMax {
			max = cMax
		}
	}

	// if max < cMax then max = cMax
	if max < cMax {
		max = cMax
	}

	return max
}
