package findLongestSubstrWithoutRepeatingChars

type LongestSubstringWithoutRepCharsFinderS1BF struct {
}

func (bf LongestSubstringWithoutRepCharsFinderS1BF) findLongestSubstringWithoutRepeatingChars(s string) int {

	var max int
	for i := 0; i < len(s); i++ {
		seenChars := make(map[uint8]bool)
		var cMax int
		for j := i; j < len(s); j++ {
			if _, ok := seenChars[s[j]]; !ok {
				cMax++
				seenChars[s[j]] = true
			} else {
				break
			}
		}

		if cMax > max {
			max = cMax
		}
	}

	return max
}
