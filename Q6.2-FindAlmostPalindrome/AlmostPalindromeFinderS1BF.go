package findPalindrome

import (
	"regexp"
	"strings"
)

type AlmostPalindromeFinderS1BF struct {
}

func (s1Bf AlmostPalindromeFinderS1BF) findAlmostPalindrome(s string) bool {
	r, _ := regexp.Compile("[^a-zA-Z0-9]")
	s = r.ReplaceAllString(s, "")
	if len(s) <= 2 {
		return true
	}

	s = strings.ToLower(s)

	// iterate over the string and remove one char at a time
	//rC := utf8.RuneCountInString(s)
	for i := range s {
		s1 := make([]rune, 0, len(s))
		for i1, c1 := range s {
			if i != i1 {
				s1 = append(s1, c1)
			}
		}
		s2 := string(s1)
		// check if the remaining chars form palindrome if yes, then return true otherwise check next
		left, right := 0, len(s2)-1
		isP := true
		for left < right {
			if s2[left] != s2[right] {
				isP = false
				break
			}
			left++
			right--
		}

		if isP {
			return true
		}
	}

	return false
}
