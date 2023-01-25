package cycledetector

import (
	"github.com/solve_common_problems_using_go"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	cNodeImpls = make([]CNode, 0)
)

type CNodeTest struct {
	testSummary string
	input       []any
	expOpt      []any
}

func testCNodeData() []CNodeTest {
	return []CNodeTest{
		{
			"normal test Data",
			[]any{1, 2, 3, 4, 5, 3},
			[]any{1, 2, 3, 4, 5, 3},
		},
		{
			"normal test Data",
			[]any{1, 2, 3, 4, 5, 3, 6},
			[]any{1, 2, 3, 4, 5, 3},
		},
	}
}

func TestCNode(t *testing.T) {
	cNodeImpls = append(cNodeImpls, CNode{})

	tests := testCNodeData()
	t.Logf("Given a list of values, find out if the singly linked list population and retrieval is working fine? [#tests=%d]", len(tests))

	for _, impl := range cNodeImpls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {

			impl.Init(tt.input)

			var pOrF string
			if assert.Equal(t, tt.expOpt, impl.AllData()) {
				pOrF = solve_common_problems_using_go.Passed
			} else {
				pOrF = solve_common_problems_using_go.Failed
			}

			t.Logf("test # %d, %s input=%v and expected output=%v, %v", i+1, pOrF, tt.input, tt.expOpt, tt.testSummary)
		}
	}
}
