package findLongestSubstr

type LongestSubstrFinderS2EffSimplified struct {
}

func (s2 LongestSubstrFinderS2EffSimplified) findLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	var longest, leftIdx int

	seenChars := make(map[rune]int)

	for rightIdx, char := range s {
		idx, ok := seenChars[char]
		if ok && idx >= leftIdx {
			leftIdx = idx + 1 // need to restart left index from the earlier position
		}
		seenChars[char] = rightIdx // update index

		d := rightIdx - leftIdx + 1
		if longest < d {
			longest = d
		}
	}

	return longest
}
