package Q4_FindTypedOutStrings

type TypedOutStringsEqualsFindersS2Efficient struct {
}

func (s2 TypedOutStringsEqualsFindersS2Efficient) findIfTypedOutStringsAreEqual(s, t string) bool {
	if len(s) == 0 && len(t) == 0 {
		return true
	}

	// iterate of both s and t at the same time from backwards, say position of s = pS = len(s) - 1 and pT = len(t) - 1
	pS, pT := len(s)-1, len(t)-1
	for pS >= 0 || pT >= 0 {
		// if any of the char is # then iterate over the loop to find out the first non hash char
		//if pS > 0 && s[pS] == hash || t[pT] == hash {
	checkHashInS:
		if pS >= 0 && s[pS] == hash {
			var skipChars int
			pS-- // since we already checked for hash
			for skipChars >= 0 && pS >= 0 {
				if s[pS] == hash {
					skipChars += 1
				} else {
					skipChars--
				}
				pS--
			}

			if pS >= 0 {
				goto checkHashInS // we need non-hash char so check for this again
			}
		}

	checkHashInT:
		if pT >= 0 && t[pT] == hash {
			var skipChars int
			pT--
			for skipChars >= 0 && pT >= 0 {
				if t[pT] == hash {
					skipChars += 1
				} else {
					skipChars--
				}
				pT--
			}

			if pT >= 0 {
				goto checkHashInT
			}
		}

		if pS >= 0 && pT >= 0 {
			if s[pS] != t[pT] { // check if they are not equal return false
				return false
			}

			// check the next number
			pS--
			pT--
		} else if pS < 0 && pT < 0 {
			break
		} else { // if there is only one of them remaining then return false
			return false
		}
	}

	return true
}
