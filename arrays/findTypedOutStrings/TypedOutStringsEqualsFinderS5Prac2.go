package findTypedOutStrings

type TypedOutStringsEqualsFinderS4Prac2 struct {
}

func (p2 TypedOutStringsEqualsFinderS4Prac2) findIfTypedOutStringsAreEqual(s, t string) bool {
	// iterate over both s and t with two pointers from backwards
	pS, pT := len(s)-1, len(t)-1
	for pS >= 0 || pT >= 0 {
		// check if s or t has hash in it
		// if anyone has it then skip the respective number of chars
	checkS:
		if pS >= 0 && s[pS] == hash {
			// say abcd#f#
			// loop till backChars is not zero
			backChars := 2
			for pS >= 0 && backChars > 0 {
				pS -= 1
				backChars -= 1
				if pS >= 0 && s[pS] == hash {
					backChars += 2
				}
			}

			goto checkS
		}

	checkT:
		if pT >= 0 && t[pT] == hash {
			backChars := 2
			for pT >= 0 && backChars > 0 {
				pT -= 1
				backChars -= 1
				if pT >= 0 && t[pT] == hash {
					backChars += 2
				}
			}

			goto checkT
		}

		if pS < 0 && pT < 0 {
			return true
		} else if pS < 0 || pT < 0 {
			return false
		} else if s[pS] != t[pT] {
			return false
		}

		pS--
		pT--
	}

	return true
}
