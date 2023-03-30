package solve_common_problems_using_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type NodeTest struct {
	testSummary string
	input       []any
	expOpt      []any
}

var (
	impls = make([]SLLNode, 0)
)

func TestNode(t *testing.T) {
	impls = append(impls, SLLNode{})

	tests := testData()
	t.Logf("Given a list of values, find out if the singly linked list population and retrieval is working fine? [#tests=%d]", len(tests))

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {

			impl.Init(tt.input)

			var pOrF string
			if assert.Equal(t, tt.expOpt, tt.input) {
				pOrF = Passed
			} else {
				pOrF = Failed
			}

			t.Logf("test # %d, %s input=%v and expected output=%v, %v", i+1, pOrF, tt.input, tt.expOpt, tt.testSummary)
		}
	}
}

func testData() []NodeTest {
	return []NodeTest{
		{
			"normal test Data",
			[]any{1, 2, 3, 4, 5},
			[]any{1, 2, 3, 4, 5},
		},
	}
}
