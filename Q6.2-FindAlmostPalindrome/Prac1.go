package findPalindrome

import (
	"regexp"
	"strings"
)

type Prac1 struct {
}

func (p1 Prac1) findAlmostPalindrome(input string) bool {
	r, _ := regexp.Compile("[^a-zA-Z0-9]")
	input = r.ReplaceAllString(input, "")
	input = strings.ToLower(input)

	pL, pR := 0, len(input)-1
	for pL < pR {
		if input[pL] != input[pR] {
			return p1.subPalindrome(pL+1, pR, input) || p1.subPalindrome(pL, pR-1, input)
		}

		pL++
		pR--
	}

	return true
}

func (p1 Prac1) subPalindrome(pL, pR int, input string) bool {
	for pL < pR {
		if input[pL] != input[pR] {
			return false
		}

		pL++
		pR--
	}

	return true
}
