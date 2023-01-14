package findPalindrome

import (
	"strings"
)

type PalindromeFinderS1BF struct {
}

func (s1 PalindromeFinderS1BF) findPalindrome(s string) bool {
	s = removePuncChars(s)
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func removePuncChars(s string) string {
	s = strings.ToLower(s)
	newS := make([]rune, 0)
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			newS = append(newS, c)
		}
	}

	return string(newS)
}
