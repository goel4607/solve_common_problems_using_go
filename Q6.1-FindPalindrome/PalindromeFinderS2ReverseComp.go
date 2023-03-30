package Q6_1_FindPalindrome

type PalindromeFinderS2ReverseComp struct {
}

func (s1 PalindromeFinderS2ReverseComp) findPalindrome(s string) bool {
	s = removePuncChars(s)
	if len(s) <= 1 {
		return true
	}

	// Q7-LinkedListReverse string
	t := make([]uint8, len(s), len(s))
	//left, right := 0, len(s)-1
	//for left < len(s) {
	//	t = append(t, s[right])
	//	left++
	//	right--
	//} // abcde t[0]=s[4]=e, t[1]=s[3]=d, t[2]=s[2]=c, t[3]=s[1]=b, t[4]=s[0]=a

	left, right := 0, len(s)-1
	for left < right {
		tmp := s[left]
		t[left] = s[right]
		t[right] = tmp
		left++
		right--
	} // abcde t[0]=s[4]=e, t[1]=s[3]=d, t[2]=s[2]=c, t[3]=s[1]=b, t[4]=s[0]=a

	if len(s)%2 != 0 { // odd number of chars
		t[left] = s[right]
	}

	// compare char by char
	for i := range s {
		if s[i] != t[i] {
			return false
		}
	}

	return true
}
