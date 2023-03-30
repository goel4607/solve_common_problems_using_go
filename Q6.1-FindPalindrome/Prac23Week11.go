package Q6_1_FindPalindrome

import (
	"regexp"
	"strings"
)

// start time: 7:00 pm
// end time: 7:06 pm
type Prac23Week11 struct {
}

func (p Prac23Week11) findPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	s = strings.ToLower(s)
	r, _ := regexp.Compile("[^a-zA-Z0-9]")
	s = r.ReplaceAllString(s, "")

	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}
