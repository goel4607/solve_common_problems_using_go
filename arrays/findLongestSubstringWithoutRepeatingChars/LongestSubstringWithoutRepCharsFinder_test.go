package findLongestSubstringWithoutRepeatingChars

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	impls = make([]FindLongestSubstringWithoutRepCharsFinder, 0)
)

func TestFindLongestSubstringWithoutRepChars(t *testing.T) {
	impls = append(impls, LongestSubstringWithoutRepCharsFinderS1BF{})

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
			t.Logf("test # %d, input=%v and expected output=%v, %v", i+1, tt.input, tt.expectedOpt, tt.info)
			actualOpt := impl.findLongestSubstringWithoutRepeatingChars(tt.input)
			assert.Equal(t, tt.expectedOpt, actualOpt)
		}
	}
}
