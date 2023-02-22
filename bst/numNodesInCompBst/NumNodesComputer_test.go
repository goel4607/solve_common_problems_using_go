package numNodesInCompBst

import (
	"github.com/solve_common_problems_using_go"
	"github.com/solve_common_problems_using_go/bst"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	impls = make([]NumNodesComputer, 0)
)

type NumNodesComputerTest struct {
	msg string
	in  []int
	out int
}

func getTests() []NumNodesComputerTest {
	return []NumNodesComputerTest{
		{
			msg: "Balanced tree with all nodes filled at each level.",
			in:  []int{20, 10, 25, 5, 12, 22, 30}, // 3 levels
			out: 7,
		},
		{
			msg: "Balanced tree with top 3 levels and only 1 node at last level.",
			in:  []int{20, 10, 25, 5, 12, 22, 30, 1}, // 3 levels complete and 1 on the last level
			out: 8,
		},
		{
			msg: "Balanced tree with top 3 levels and only 1 node at last level.",
			in:  []int{20, 10, 25, 5, 12, 22, 30, 1, 4, 11, 14, 21}, // 3 levels complete and 5 on the last level
			out: 12,
		},
	}
}

func TestTypedOutStringsEqualFinder(t *testing.T) {
	impls = append(
		impls,
		//RSVSolutionUsingBFS{},
		//RSVSolutionUsingDFS{},
	)

	tests := getTests()
	bytes, err := os.ReadFile("problem.md")
	assert.Nil(t, err)
	t.Logf("Problem: %s [#tests=%d]", bytes, len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")
	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {

			b := &bst.Node{Data: tt.in[0]}
			for j, d := range tt.in {
				if j != 0 {
					b.Insert(d)
				}
			}

			actualOpt := impl.computeNumNodes(b)
			assert.Equal(t, tt.out, actualOpt)

			var pOrF string
			if assert.Equal(t, tt.out, actualOpt) {
				pOrF = solve_common_problems_using_go.Passed
			} else {
				pOrF = solve_common_problems_using_go.Failed
			}

			t.Logf("test # %d, %s input=%v and expected output=%v, %v", i+1, pOrF, tt.in, tt.out, tt.msg)
		}
	}
}
