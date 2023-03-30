package findPalindrome

import (
	"regexp"
	"strings"
)

type AlmostPalindromeFinderS3EfficientBetterReadability struct {
}

func (s3 AlmostPalindromeFinderS3EfficientBetterReadability) findAlmostPalindrome(s string) bool {
	r, _ := regexp.Compile("[^a-zA-Z0-9]")
	s = r.ReplaceAllString(s, "")
	s = strings.ToLower(s)

	if len(s) <= 2 {
		return true
	}

	lt, rt := 0, len(s)-1
	for lt < rt {
		if s[lt] != s[rt] {
			return s3.findSubPalindrome(s, lt+1, rt) || s3.findSubPalindrome(s, lt, rt-1)
		}
		lt++
		rt--
	}

	return false
}

func (s3 AlmostPalindromeFinderS3EfficientBetterReadability) findSubPalindrome(s string, lt, rt int) bool {
	for lt < rt {
		if s[lt] != s[rt] {
			return false
		}

		lt++
		rt--
	}

	return true
}
