package Q15_MaxDepthOfBinaryTree

import (
	"github.com/solve_common_problems_using_go"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	impls = make([]MaxDepthFinder, 0)
)

type MaxDepthFinderTest struct {
	msg string
	in  []int
	out int
}

func getTests() []MaxDepthFinderTest {
	return []MaxDepthFinderTest{
		{
			msg: "Tree is almost flat i.e. only one node at each level.",
			in:  []int{1},
			out: 1,
		},
		{
			msg: "Tree is almost flat i.e. only one node at each level.",
			in:  []int{1, 2, 3, 4, 5},
			out: 5,
		},
		{
			msg: "Tree is almost flat i.e. only one node at each level.",
			in:  []int{20, 10, 5, 6, 7, 8, 9},
			out: 7,
		},
		{
			msg: "Balanced tree with all nodes filled at each level.",
			in:  []int{20, 10, 25, 5, 12, 22, 30},
			out: 3,
		},
		{
			msg: "Balanced tree with top 3 levels and only 1 node at last 2 levels.",
			in:  []int{20, 10, 25, 5, 12, 22, 30, 1, 2},
			out: 5,
		},
	}
}

func TestTypedOutStringsEqualFinder(t *testing.T) {
	impls = append(
		impls,
		//S1MaxDepthFinder{},
		Prac23Week11UsingDFS{},
	)

	assert.NotEmptyf(t, impls, "no implemantations present!!")

	tests := getTests()
	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {

			b := &solve_common_problems_using_go.BSTNode{Data: tt.in[0]}
			for j, d := range tt.in {
				if j != 0 {
					b.Insert(d)
				}
			}

			actualOpt := impl.findMaxDepth(b)
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
