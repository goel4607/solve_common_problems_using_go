package Q22_WallsAndGates

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
	msg string
	in  [][]int
	out [][]int
}

func getTests() []InterfaceTest {
	return []InterfaceTest{
		{
			msg: "4 by 4 array with 2 gates and 5 walls and no un-reachable value(s).",
			in: [][]int{
				{INF, WALL, GATE, INF},
				{INF, INF, INF, WALL},
				{INF, WALL, INF, WALL},
				{GATE, WALL, INF, INF},
			},
			out: [][]int{
				{3, WALL, GATE, 1},
				{2, 2, 1, WALL},
				{1, WALL, 2, WALL},
				{GATE, WALL, 3, 4},
			},
		},
		{
			msg: "4 by 4 array with 2 gates and 5 walls with one un-reachable value.",
			in: [][]int{
				{INF, WALL, GATE, INF},
				{WALL, INF, INF, WALL},
				{INF, WALL, INF, WALL},
				{GATE, WALL, INF, INF},
			},
			out: [][]int{
				{INF, WALL, GATE, 1},
				{WALL, 2, 1, WALL},
				{1, WALL, 2, WALL},
				{GATE, WALL, 3, 4},
			},
		},
		{
			msg: "empty array",
			in: [][]int{
				{},
				{},
			},
			out: [][]int{
				{},
				{},
			},
		},
	}
}

func TestInterface(t *testing.T) {
	impls = append(
		impls,
		//Soln1{},
		//Prac23AprNum1BFS{},
		//Prac23AprNum2DFS{},
		Prac23June01DFS{},
	)

	tests := getTests()
	bytes, err := os.ReadFile("problem.md")
	assert.Nil(t, err)
	t.Logf("Problem: \n%s \n[#tests=%d]", bytes, len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")
	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {

			impl.FillRoomsWithNumHopsToNearestGate(tt.in)

			var pOrF string
			if assert.Equal(t, tt.out, tt.in) {
				pOrF = solve_common_problems_using_go.Passed
			} else {
				pOrF = solve_common_problems_using_go.Failed
			}

			t.Logf("test # %d, %s input=%v and expected output=%v, %v", i+1, pOrF, tt.in, tt.out, tt.msg)
		}
	}
}
