package findPalindrome

import (
	"github.com/solve_common_problems_using_go"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	impls = make([]PalindromeFinder, 0)
)

func TestPalindrome(t *testing.T) {
	impls = append(impls,
		//PalindromeFinderS1LeftAndRightPts{},
		//PalindromeFinderS2ReverseComp{},
		//PalindromeFinderS3CenterExpand{},
		Prac23Week11{},
	)

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		td := getTestData()
		t.Logf("Given a string, find out whether it is a palindrome? [#tests=%d]", len(td))

		for i, tt := range td {

			actualOpt := impl.findPalindrome(tt.input)
			var pOrF string
			if assert.Equal(t, tt.expectedOpt, actualOpt) {
				pOrF = solve_common_problems_using_go.Passed
			} else {
				pOrF = solve_common_problems_using_go.Failed
			}

			t.Logf("test # %d, %s input=%v and expected output=%v, %v", i+1, pOrF, tt.input, tt.expectedOpt, tt.testSummary)
		}
	}
}

type PalindromeTester struct {
	testSummary string
	input       string
	expectedOpt bool
}

func getTestData() []PalindromeTester {
	return []PalindromeTester{
		{"+ve test, a long one", "A man, a plan, a canal: Panama", true},
		{"+ve test, a simple one", "aba", true},
		{"-ve test, a simple one", "abc", false},
		{"+ve test, only one char", "a", true},
		{"+ve test, two same chars", "aa", true},
		{"-ve test, empty string", "", true},
		{"+ve test, two words", "race car", true},
		{"+ve test, simple odd number of chars", "aabaa", true},
		{"+ve test, simple even number of chars", "abba", true},
	}
}
