package minTimeToEachNode

import (
	"github.com/solve_common_problems_using_go"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	impls = make([]Interface, 0)
)

type InterfaceTest struct {
	msg          string
	n            int
	startingNode int
	eachNodeTime []Node
	out          int
}

func getTests() []InterfaceTest {
	return []InterfaceTest{
		{
			msg:          "+ve, all reachable nodes.",
			n:            5,
			startingNode: 0,
			eachNodeTime: []Node{
				0: {0, 1, 9}, // 1 ---9---> 2
				1: {0, 3, 2},
				2: {1, 4, 1},
				3: {3, 1, 4},
				4: {3, 4, 6},
				5: {2, 1, 3},
				6: {4, 2, 7},
				7: {2, 0, 5},
			},
			out: 14,
		},
		{
			msg:          "-ve, node 1 is not reachable.",
			n:            2,
			startingNode: 0,
			eachNodeTime: []Node{
				0: {1, 2, 4}, //2----3---->4
			},
			out: -1,
		},
		{
			msg:          "-ve, there is no path from 1 to 3.",
			n:            3,
			startingNode: 0,
			eachNodeTime: []Node{
				0: {0, 1, 8}, //1----8---->4
				1: {2, 0, 3}, //3----3---->1
			},
			out: -1,
		},
	}
}

func TestInterface(t *testing.T) {
	impls = append(
		impls,
		Soln1{},
	)

	tests := getTests()
	bytes, err := os.ReadFile("problem.md")
	assert.Nil(t, err)
	t.Logf("\nProblem: %s \n[#tests=%d]", bytes, len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")
	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {

			actual := impl.findMinTime(tt.n, tt.startingNode, tt.eachNodeTime)

			var pOrF string
			if assert.Equal(t, tt.out, actual) {
				pOrF = solve_common_problems_using_go.Passed
			} else {
				pOrF = solve_common_problems_using_go.Failed
			}

			t.Logf("test # %d, %s input(s): [n=%v, pre-reqs=%v] and expected output=%v, %v", i+1, pOrF, tt.n, tt.eachNodeTime, tt.out, tt.msg)
		}
	}
}
