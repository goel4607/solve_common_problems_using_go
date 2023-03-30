package Q5_FindLongestSubstr

type LongestSubstrFinderS5Prac2 struct {
}

func (p LongestSubstrFinderS5Prac2) findLongestSubstring(s string) int {
	idxByChar := make(map[rune]int)

	var pL, pR, max int
	for i, v := range s {
		// check if the char is read first time or not, if not then inc the max and store the element in the map
		if oldPos, ok := idxByChar[v]; ok {
			if oldPos >= pL {
				pL = oldPos + 1 // reset the left pointer
			}
		}

		idxByChar[v] = i

		cMax := pR - pL + 1
		if max < cMax {
			max = cMax
		}

		pR++
	}

	return max
}
