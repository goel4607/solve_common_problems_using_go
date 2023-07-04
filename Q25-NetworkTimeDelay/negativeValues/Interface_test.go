package negativeValues

import (
	"github.com/solve_common_problems_using_go"
	Q25_NetworkTimeDelay "github.com/solve_common_problems_using_go/Q25-NetworkTimeDelay"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	impls = make([]Q25_NetworkTimeDelay.Interface, 0)
)

type InterfaceTest struct {
	msg          string
	n            int
	startingNode int
	eachNodeTime [][]int
	out          int
}

func getTests() []InterfaceTest {
	return []InterfaceTest{
		{
			msg:          "+ve, all reachable nodes.",
			n:            5,
			startingNode: 1,
			eachNodeTime: [][]int{
				0: {1, 2, 9}, // 1 ---9---> 2
				1: {1, 4, 2},
				2: {2, 5, -3},
				3: {4, 2, -4},
				4: {4, 5, 6},
				5: {3, 2, 3},
				6: {5, 3, 7},
				7: {3, 1, 5},
			},
			out: 2,
		},
		{
			msg:          "+ve, all reachable nodes.",
			n:            5,
			startingNode: 1,
			eachNodeTime: [][]int{
				0: {1, 2, 9}, // 1 ---9---> 2
				1: {1, 4, 2},
				2: {2, 5, 1},
				3: {4, 2, 4},
				4: {4, 5, 6},
				5: {3, 2, 3},
				6: {5, 3, 7},
				7: {3, 1, 5},
			},
			out: 14,
		},
		{
			msg:          "-ve, node 1 is not reachable.",
			n:            2,
			startingNode: 1,
			eachNodeTime: [][]int{
				0: {2, 3, 4}, //2----3---->4
			},
			out: -1,
		},
		{
			msg:          "-ve, there is no path from 1 to 3.",
			n:            3,
			startingNode: 1,
			eachNodeTime: [][]int{
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
		//Soln1{},
		//Prac1{},
		//Prac23AprDijk{},
		//Prac23MayBellmanFord{},
		//Prac23June01BF{},
		//Prac23June02BF{},
		Prac23July01BellmanFord{},
	)

	tests := getTests()
	bytes, err := os.ReadFile("problem.md")
	assert.Nil(t, err)
	t.Logf("\nProblem: %s \n[#tests=%d]", bytes, len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")
	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {

			actual := impl.NetworkDelayTime(tt.eachNodeTime, tt.n, tt.startingNode)

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
