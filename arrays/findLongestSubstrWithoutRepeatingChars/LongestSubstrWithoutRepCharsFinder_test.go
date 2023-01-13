package findLongestSubstrWithoutRepeatingChars

import (
	"github.com/solve_common_problems_using_go"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	impls = make([]FindLongestSubstringWithoutRepCharsFinder, 0)
)

func TestFindLongestSubstringWithoutRepChars(t *testing.T) {
	impls = append(impls, LongestSubstringWithoutRepCharsFinderS1BF{}, LongestSubstrWithoutRepeatingCharsS2Eff{}, LongestSubstrWithoutRepeatingCharsS2EffSimplified{})

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		inputs := []struct {
			info        string
			input       string
			expectedOpt int
		}{
			{
				"A positive test case with greatest value is among the area walls",
				"abccabb",
				3,
			},
			{
				"A positive test case with greatest value is among the area walls",
				"aaaaaa",
				1,
			},
			{
				"Another positive test case with middle value greatest but it is not amongst the wall",
				"",
				0,
			},
			{
				"Another positive test case when diff is only one",
				"abcbda",
				4,
			},
		}

		for i, tt := range inputs {

			actualOpt := impl.findLongestSubstringWithoutRepeatingChars(tt.input)
			var pOrF string
			if assert.Equal(t, tt.expectedOpt, actualOpt) {
				pOrF = solve_common_problems_using_go.Passed
			} else {
				pOrF = solve_common_problems_using_go.Failed
			}

			t.Logf("test # %d, %s input=%v and expected output=%v, %v", i+1, pOrF, tt.input, tt.expectedOpt, tt.info)
		}
	}
}
