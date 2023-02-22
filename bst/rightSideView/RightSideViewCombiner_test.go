package rightSideView

import (
	"github.com/solve_common_problems_using_go"
	"github.com/solve_common_problems_using_go/bst"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	impls = make([]RightSideViewCombiner, 0)
)

type LevelValuesCombinerTest struct {
	msg string
	in  []int
	out []int
}

func getTests() []LevelValuesCombinerTest {
	return []LevelValuesCombinerTest{
		{
			msg: "Tree is almost flat i.e. only one node at each level.",
			in:  []int{1, 2, 3, 4, 5},
			out: []int{1, 2, 3, 4, 5},
		},
		{
			msg: "Tree is almost flat i.e. only one node at each level.",
			in:  []int{20, 10, 5, 6, 7, 8, 9},
			out: []int{20, 10, 5, 6, 7, 8, 9},
		},
		{
			msg: "Balanced tree with all nodes filled at each level.",
			in:  []int{20, 10, 25, 5, 12, 22, 30},
			out: []int{20, 25, 30},
		},
		{
			msg: "Balanced tree with top 3 levels and only 1 node at last 2 levels.",
			in:  []int{20, 10, 25, 5, 12, 22, 30, 1, 2},
			out: []int{20, 25, 30, 1, 2},
		},
	}
}

func TestTypedOutStringsEqualFinder(t *testing.T) {
	impls = append(
		impls,
		RSVSolutionUsingBFS{},
		RSVSolutionUsingDFS{},
	)

	tests := getTests()
	bytes, err := os.ReadFile("problem.md")
	assert.Nil(t, err)
	t.Logf("Problem: %s? [#tests=%d]", bytes, len(tests))

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

			actualOpt := impl.combineRightSideView(b)
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
