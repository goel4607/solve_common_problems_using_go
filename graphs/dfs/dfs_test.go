package dfs

import (
	"github.com/solve_common_problems_using_go"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	impls = make([]DFS, 0)
)

type InterfaceTest struct {
	msg string
	in  [][]int
	out []int
}

func getTests() []InterfaceTest {
	return []InterfaceTest{
		{
			msg: "a simple graph adjacency list, where each sub-array represents a list of nodes reachable at that index",
			in: [][]int{
				0: {1, 3},
				1: {0},
				2: {3, 8},
				3: {4, 5, 2, 0},
				4: {6, 3},
				5: {3},
				6: {7, 4},
				7: {6},
				8: {2},
			},
			out: []int{0, 1, 3, 4, 6, 7, 5, 2, 8},
		},
	}
}

func TestInterface(t *testing.T) {
	impls = append(
		impls,
		DFSSoln1{},
	)

	tests := getTests()
	bytes, err := os.ReadFile("problem.md")
	assert.Nil(t, err)
	t.Logf("Problem: %s [#tests=%d]", bytes, len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")
	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {

			a := impl.dfs(tt.in)

			var pOrF string
			if assert.Equal(t, tt.out, a) {
				pOrF = solve_common_problems_using_go.Passed
			} else {
				pOrF = solve_common_problems_using_go.Failed
			}

			t.Logf("test # %d, %s input=%v and expected output=%v, %v", i+1, pOrF, tt.in, tt.out, tt.msg)
		}
	}
}
