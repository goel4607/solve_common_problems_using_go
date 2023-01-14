package findPalindrome

type PalindromeFinderS3CenterExpand struct {
}

func (s3 PalindromeFinderS3CenterExpand) findPalindrome(s string) bool {
	s = removePuncChars(s)
	if len(s) <= 1 {
		return true
	}

	c := len(s) / 2
	for i := 0; i < c; i++ {
		left := c - i - 1
		if left < 0 {
			break
		}

		right := c + i
		if len(s)%2 != 0 {
			right = c + i + 1 // compensate for rounding
		}
		if right > len(s)-1 {
			break
		}

		if s[left] != s[right] {
			return false
		}
	}

	return true
}
