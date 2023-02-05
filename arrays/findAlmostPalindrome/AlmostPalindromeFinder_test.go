package findPalindrome

import (
	"github.com/solve_common_problems_using_go"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	impls = make([]AlmostPalindromeFinder, 0)
)

type AlmostPalindromeTester struct {
	testSummary string
	input       string
	expectedOpt bool
}

func getTestData() []AlmostPalindromeTester {
	return []AlmostPalindromeTester{
		//{"+ve test, a long one", "A man, a plan, a canal: Panama", true},
		{"+ve test, a long almost one", "A man, a plan, a canal: anama", true},
		{"+ve test, a long almost one", "A an, a plan, a canal: Panama", true},
		{"-ve test, a long almost but two chars, so not a palindrome", "A n, a plan, a canal: Panama", false},
		//{"+ve test, a simple one", "aba", true},
		{"+ve test, a simple almost even one", "abaa", true},
		{"+ve test, a simple almost one", "aaba", true},
		{"+ve test, a simple almost odd one", "aaaba", true},
		{"+ve test, a simple almost odd one", "aaabba", false},
		{"-ve test, a simple one", "abc", false},
		{"-ve test, a simple almost one", "acbc", true},
		{"+ve test, only one char", "a", true},
		{"+ve test, only one almost with two chars", "ab", true},
		{"+ve test, two same chars", "aa", true},
		{"-ve test, empty string", "", true},
		//{"+ve test, two words", "race car", true},
		//{"+ve test, simple odd number of chars", "aabaa", true},
		//{"+ve test, simple even number of chars", "abba", true},
	}
}
func TestPalindrome(t *testing.T) {
	impls = append(impls,
		AlmostPalindromeFinderS1BF{},
		AlmostPalindromeFinderS2Eff{},
		AlmostPalindromeFinderS3EfficientBetterReadability{},
		Prac1{},
	)

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		td := getTestData()
		t.Logf("Given a string, find out whether it is an almost palindrome (removing a single char makes it a palindrome)? [#tests=%d]", len(td))

		for i, tt := range td {

			actualOpt := impl.findAlmostPalindrome(tt.input)
			var pOrF string
			if assert.Equal(t, tt.expectedOpt, actualOpt) {
				pOrF = solve_common_problems_using_go.Passed
			} else {
				pOrF = solve_common_problems_using_go.Failed
			}

			t.Logf("test # %2d, %s input=%q and expected output=%v, %v", i+1, pOrF, tt.input, tt.expectedOpt, tt.testSummary)
		}
	}
}
