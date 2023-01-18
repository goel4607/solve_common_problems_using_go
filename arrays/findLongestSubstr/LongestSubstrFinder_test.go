package findLongestSubstr

import (
	"github.com/solve_common_problems_using_go"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	impls = make([]FindLongestSubstrFinder, 0)
)

type LongestSubstrFinderTest struct {
	info        string
	input       string
	expectedOpt int
}

func testData() []LongestSubstrFinderTest {
	return []LongestSubstrFinderTest{
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
		{
			"Another positive test case when diff is only one",
			"abcbdabcd",
			4,
		},
	}
}

func TestFindLongestSubstringWithoutRepChars(t *testing.T) {
	impls = append(
		impls,
		//LongestSubstrFinderS1BF{},
		//LongestSubstrFinderS2Eff{},
		//LongestSubstrFinderS2EffSimplified{},
		LongestSubstrFinderS4Prac1{},
	)

	tests := testData()

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {

			actualOpt := impl.findLongestSubstring(tt.input)
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
