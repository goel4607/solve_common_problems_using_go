package mnReverser

import (
	"github.com/solve_common_problems_using_go"
	"github.com/solve_common_problems_using_go/linkedlist/singlylinkedlist"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MNReverseTest struct {
	testSummary string
	data        []any
	m           int
	n           int
	expOpt      []any
}

func testData() []MNReverseTest {
	return []MNReverseTest{
		{
			"normal test data with m and m in-between values",
			[]any{1, 2, 3, 4, 5, 6, 7}, 3, 5,
			[]any{1, 2, 5, 4, 3, 6, 7},
		},
		{
			"normal test data with n as the last value",
			[]any{1, 2, 3, 4, 5, 6, 7}, 3, 7,
			[]any{1, 2, 7, 6, 5, 4, 3},
		},
		{
			"normal test data with m as the first value",
			[]any{1, 2, 3, 4, 5, 6, 7}, 1, 5,
			[]any{5, 4, 3, 2, 1, 6, 7},
		},
		{
			"normal test data with m as the first value",
			[]any{1, 2, 3, 4, 5, 6, 7}, 1, 7,
			[]any{7, 6, 5, 4, 3, 2, 1},
		},
	}
}

var (
	impls = make([]MNReverser, 0)
)

func TestNode(t *testing.T) {
	impls =
		append(impls,
			&MNReverserS1BF{},
			//&ReverserS2Repeat{},
		)

	tests := testData()
	t.Logf("Given a list of values along with starting and ending positions m and n, reverse the list such that only the elements between m & n are reversed [#tests=%d]", len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")
	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			headNode := &singlylinkedlist.Node{}
			headNode.Init(tt.data)

			rNode := impl.reverse(headNode, tt.m, tt.n)
			var pOrF string
			if assert.Equal(t, tt.expOpt, rNode.AllData()) {
				pOrF = solve_common_problems_using_go.Passed
			} else {
				pOrF = solve_common_problems_using_go.Failed
			}

			t.Logf("test # %d, %s input[data=%v, m=%v, n=%v] and expected output=%v, %v", i+1, pOrF, tt.data, tt.m, tt.n, tt.expOpt, tt.testSummary)
		}
	}
}
