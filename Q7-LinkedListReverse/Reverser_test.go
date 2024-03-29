package Q7_LinkedListReverse

import (
	"github.com/solve_common_problems_using_go"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ReverseTest struct {
	testSummary string
	input       []any
	expOpt      []any
}

var (
	impls = make([]Reverser, 0)
)

func TestReverserImplementations(t *testing.T) {
	impls =
		append(impls,
			//&ReverserS1BF{},
			//&ReverserS2Repeat{},
			&Prac23Week11{},
		)

	tests := testData()
	t.Logf("Given a list of values, find out if the singly linked list population and retrieval is working fine? [#tests=%d]", len(tests))

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			n := &solve_common_problems_using_go.SLLNode{}
			n.Init(tt.input)

			rNode := impl.reverse(n)
			var pOrF string
			if assert.Equal(t, tt.expOpt, rNode.AllData()) {
				pOrF = solve_common_problems_using_go.Passed
			} else {
				pOrF = solve_common_problems_using_go.Failed
			}

			t.Logf("test # %d, %s input=%v and expected output=%v, %v", i+1, pOrF, tt.input, tt.expOpt, tt.testSummary)
		}
	}
}

func testData() []ReverseTest {
	return []ReverseTest{
		{
			"normal test data",
			[]any{1, 2, 3, 4, 5},
			[]any{5, 4, 3, 2, 1},
		},
	}
}
