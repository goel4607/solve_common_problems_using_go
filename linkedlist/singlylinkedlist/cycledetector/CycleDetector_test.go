package cycledetector

import (
	"github.com/solve_common_problems_using_go"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	cycleDetectorImpls = make([]CycleDetector, 0)
)

type CycleDetectorTest struct {
	testSummary string
	data        []any
	expOpt      bool
}

func testData() []CycleDetectorTest {
	return []CycleDetectorTest{
		{
			"normal test data with m and m in-between values",
			[]any{1, 2, 3, 4, 5, 6, 7, 3, 7},
			true,
		},
		{
			"normal test data with m and m in-between values",
			[]any{1, 2},
			false,
		},
		{
			"normal test data with m and m in-between values",
			[]any{1},
			false,
		},
		{
			"normal test data with n as the last value",
			[]any{1, 2, 3, 4, 5, 6, 7},
			false,
		},
		{
			"normal test data with m as the first value",
			[]any{1, 2, 3, 4, 5, 3, 7},
			true,
		},
	}
}

func TestMNReverserImplementations(t *testing.T) {
	cycleDetectorImpls = append(
		cycleDetectorImpls,
		//Solution1BruteForce{},
		//Solution2FastSlowPtrs{},
		Prac23Week11{},
	)

	tests := testData()
	t.Logf("Given a list of values along with starting and ending positions m and n, reverse the list such that only the elements between m & n are reversed [#tests=%d]", len(tests))

	assert.NotEmptyf(t, cycleDetectorImpls, "no implemantations present!!")
	for _, impl := range cycleDetectorImpls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			headNode := &CNode{}
			headNode.Init(tt.data)

			rNode := impl.detectCycle(headNode)
			var pOrF string
			if assert.Equal(t, tt.expOpt, rNode) {
				pOrF = solve_common_problems_using_go.Passed
			} else {
				pOrF = solve_common_problems_using_go.Failed
			}

			t.Logf("test # %d, %s input: data=%v and expected output=%v, %v", i+1, pOrF, tt.data, tt.expOpt, tt.testSummary)
		}
	}
}
