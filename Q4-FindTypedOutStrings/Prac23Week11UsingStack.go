package Q4_FindTypedOutStrings

//start time: 8:58 pm
//end time: 9:35 pm

type Prac23Week11 struct {
}

func (p Prac23Week11) findIfTypedOutStringsAreEqual(s, t string) bool {
	//declare pS and pT as pointers to last element of s and t
	pS, pT := len(s)-1, len(t)-1

	//loop till either pS is >= 0 or pT is >= 0
	for pS >= 0 || pT >= 0 {
		//checkForSHash:
		if pS >= 0 && s[pS] == hash {
			backChars := 2
			for pS >= 0 && backChars > 0 {
				backChars--
				pS--
				if pS >= 0 && s[pS] == hash {
					backChars += 2
				}
			}
			//goto checkForSHash
		}

		//checkForTHash:
		if pT >= 0 && t[pT] == hash {
			backChars := 2
			for pT >= 0 && backChars > 0 {
				backChars--
				pT--
				if pT >= 0 && t[pT] == hash {
					backChars += 2
				}
			}

			//goto checkForTHash
		}

		if pS >= 0 && pT >= 0 {
			if s[pS] != t[pT] {
				return false
			}
		} else if pS < 0 && pT < 0 {
			return true
		} else {
			return false
		}

		pS--
		pT--
	}

	return true
}
