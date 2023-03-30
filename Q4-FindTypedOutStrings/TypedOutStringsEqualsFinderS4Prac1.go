package Q4_FindTypedOutStrings

type TypedOutStringsEqualFinderS4Prac1 struct {
}

func (s4 TypedOutStringsEqualFinderS4Prac1) findIfTypedOutStringsAreEqual(s, t string) bool {
	// init i and j to the last char of s and t
	i, j := len(s)-1, len(t)-1

	// iterate over s and t from the last char to 0
	for i >= 0 || j >= 0 {

		//checkForHashInS:
		// if s[i] is hash then back-track to the non-hash char
		if i >= 0 && s[i] == hash {
			back := 2
			for i >= 0 && back > 0 {
				i--
				back--

				if i >= 0 && s[i] == hash {
					back += 2
				}
			}

			//goto checkForHashInS
		}

		//checkForHashInT:
		// if s[j] is hash then back-track to the non-hash char
		if j >= 0 && t[j] == hash {
			back := 2
			for j >= 0 && back > 0 {
				j--
				back--

				if j >= 0 && t[j] == hash {
					back += 2
				}
			}

			//goto checkForHashInT
		}

		// if s[i] and t[j] are both non hash
		if i >= 0 && j >= 0 {
			// check if they are not equal then return false otherwise dec i and j
			if s[i] != t[j] {
				return false
			} else {
				i--
				j--
			}
		} else if i < 0 && j < 0 {
			return true // both strings are fully checked, so all match
		} else {
			return false // only one of them is rem so non-match
		}
	}

	return true
}
