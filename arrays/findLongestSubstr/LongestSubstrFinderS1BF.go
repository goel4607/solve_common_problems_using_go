package findLongestSubstr

type LongestSubstrFinderS1BF struct {
}

func (bf LongestSubstrFinderS1BF) findLongestSubstring(s string) int {

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
