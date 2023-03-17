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
	numNodes     int
	startingNode int
	edges        []Node
	out          int
}

func getTests() []InterfaceTest {
	return []InterfaceTest{
		{
			msg:          "+ve, all reachable nodes.",
			numNodes:     5,
			startingNode: 1,
			edges: []Node{
				0: {1, 2, 9}, // 1 ---9---> 2
				1: {3, 2, 3},
				2: {5, 3, 7},
				3: {3, 1, 5},
				4: {2, 5, -3},
				5: {4, 5, 6},
				6: {1, 4, 2},
				7: {4, 2, -4},
			},
			out: 2,
		},
		{
			msg:          "-ve, node 1 is not reachable.",
			numNodes:     2,
			startingNode: 1,
			edges: []Node{
				0: {1, 2, 4}, //2----3---->4
			},
			out: 4,
		},
		{
			msg:          "-ve, there is no path from 1 to 3.",
			numNodes:     3,
			startingNode: 1,
			edges: []Node{
				0: {1, 2, 8}, //1----8---->4
				1: {3, 1, 3}, //3----3---->1
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

			actual := impl.findMinTime(tt.numNodes, tt.edges, tt.startingNode)

			var pOrF string
			if assert.Equal(t, tt.out, actual) {
				pOrF = solve_common_problems_using_go.Passed
			} else {
				pOrF = solve_common_problems_using_go.Failed
			}

			t.Logf("test # %d, %s input(s): [numNodes=%v, pre-reqs=%v] and expected output=%v, %v", i+1, pOrF, tt.numNodes, tt.edges, tt.out, tt.msg)
		}
	}
}
