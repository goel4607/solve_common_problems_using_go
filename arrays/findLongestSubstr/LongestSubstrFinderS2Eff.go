package findLongestSubstr

type LongestSubstrFinderS2Eff struct {
}

func (s2 LongestSubstrFinderS2Eff) findLongestSubstring(s string) int {

	var longest, cMax, leftIdx int

	seenChars := make(map[rune]int)

	for r, c := range s {
		if idx, ok := seenChars[c]; !ok {
			cMax++
		} else if idx >= leftIdx {
			if longest < cMax {
				longest = cMax
			}
			leftIdx = idx + 1
			cMax = r - leftIdx + 1
		} else {
			cMax++
		}
		// store the index if not found
		seenChars[c] = r
	}

	if longest < cMax {
		longest = cMax
	}

	return longest
}
