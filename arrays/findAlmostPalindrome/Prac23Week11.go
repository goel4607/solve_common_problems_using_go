package findPalindrome

import (
	"regexp"
	"strings"
)

type Prac23Week11 struct {
}

func (p Prac23Week11) findAlmostPalindrome(s string) bool {
	s = strings.ToLower(s)
	r, _ := regexp.Compile("[^a-zA-Z]")
	s = r.ReplaceAllString(s, "")

	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return p.findAlmostPalindromeAtIndexes(s, i+1, j) || p.findAlmostPalindromeAtIndexes(s, i, j-1)
		}
		i++
		j--
	}

	return true
}

func (p Prac23Week11) findAlmostPalindromeAtIndexes(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
