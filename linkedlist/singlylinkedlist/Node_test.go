package singlylinkedlist

import (
	"github.com/solve_common_problems_using_go"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	impls = make([]Node, 0)
)

func TestNode(t *testing.T) {
	impls = append(impls, Node{})

	tests := testData()
	t.Logf("Given a list of values, find out if the singly linked list population and retrieval is working fine? [#tests=%d]", len(tests))

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {

			impl.Init(tt.input)

			var pOrF string
			if assert.Equal(t, tt.expOpt, tt.input) {
				pOrF = solve_common_problems_using_go.Passed
			} else {
				pOrF = solve_common_problems_using_go.Failed
			}

			t.Logf("test # %d, %s input=%v and expected output=%v, %v", i+1, pOrF, tt.input, tt.expOpt, tt.testSummary)
		}
	}
}

type NodeTest struct {
	testSummary string
	input       []any
	expOpt      []any
}

func testData() []NodeTest {
	return []NodeTest{
		{
			"normal test data",
			[]any{1, 2, 3, 4, 5},
			[]any{1, 2, 3, 4, 5},
		},
	}
}
