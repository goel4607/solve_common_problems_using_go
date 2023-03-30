package findPalindrome

import (
	"regexp"
	"strings"
)

type AlmostPalindromeFinderS2Eff struct {
}

func (s2 AlmostPalindromeFinderS2Eff) findAlmostPalindrome(s string) bool {
	r, _ := regexp.Compile("[^a-zA-Z0-9]")
	s = r.ReplaceAllString(s, "")

	if len(s) <= 2 {
		return true
	}

	s = strings.ToLower(s)

	var checkedSkipOneLeft, checkedSkipOneRight bool

	// iterate over the loop keeping left and right index pointers
	left, right := 0, len(s)-1
	for left < right {
		// compare if left != right
		if s[left] != s[right] {
			// check checkedSkipOneONLeft
			if !checkedSkipOneLeft {
				// if not then set checkedSkipOneONLeft = true and inc left,
				checkedSkipOneLeft = true
				left++
			} else if !checkedSkipOneRight { // check checkedSkipOneRight
				// if not then set it to true and then inc right
				left--
				checkedSkipOneRight = true
				right--
				// if both true
				// then return false
			} else {
				return false
			}
		} else {
			left++
			right--
		}
	}
	// if anyone of the above two variables or both are tre then return true
	// return false
	if checkedSkipOneLeft || checkedSkipOneRight {
		return true
	}

	return false
}
