package Q24_Course_scheduler

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
	msg     string
	n       int
	preReqs [][]int
	out     bool
}

//
//								4
//								5
//						3		2
//							0
// 							1

func getTests() []InterfaceTest {
	return []InterfaceTest{
		{
			msg: "+ve, course dependency with no cycle.",
			n:   7,
			preReqs: [][]int{
				0: {1, 0},
				1: {2, 1},
				2: {2, 5},
				3: {0, 3},
				4: {3, 5},
				5: {4, 3},
				6: {4, 5},
			},
			out: true,
		},
		{
			msg: "+ve, course dependency with no cycles.",
			n:   7,
			preReqs: [][]int{
				{1, 0},
				{0, 3},
				{0, 2},
				{3, 2},
				{2, 5},
				{4, 5},
				{5, 6},
				{2, 4},
			},
			out: true,
		},
		{
			msg: "+ve, course dependency with one cycle.",
			n:   7,
			preReqs: [][]int{
				0: {1, 0},
				1: {2, 1},
				2: {5, 2},
				3: {0, 3},
				4: {3, 5},
				5: {4, 3},
				6: {4, 5},
			},
			out: false,
		},
		{
			msg: "-ve, course dependency with one cycle.",
			n:   7,
			preReqs: [][]int{
				0: {1, 0},
				1: {2, 1},
				2: {2, 5},
				3: {0, 3},
				4: {3, 5},
				5: {4, 3},
				6: {5, 4},
			},
			out: false,
		},
		{
			msg: "-ve, course dependency with two cycles.",
			n:   7,
			preReqs: [][]int{
				0: {1, 0},
				1: {2, 1},
				2: {5, 2},
				3: {0, 3},
				4: {3, 5},
				5: {4, 3},
				6: {5, 4},
			},
			out: false,
		},
		{
			msg: "-ve, course dependency with two cycles.",
			n:   2,
			preReqs: [][]int{
				0: {1, 0},
			},
			out: true,
		},
		{
			msg:     "-ve, course dependency with two cycles.",
			n:       0,
			preReqs: [][]int{},
			out:     true,
		},
	}
}

func TestInterface(t *testing.T) {
	impls = append(
		impls,
		//Soln1BruteForce{},
		//Soln2TopSort{},
		//Soln3TopSortEfficient{},
		//Prac23AprTopSort{},
		//Prac230601BFS{},
		//Prac230601TopSort{},
		//Prac230602BFS{},
		//Prac23June02TopSort{},
		//Prac23June03TopSort{},
		//Prac23June03BFS{},
		//Prac23June04TopSort{},
		//Prac23July01TopSort{},
		Prac23July01BFS{},
	)

	tests := getTests()
	bytes, err := os.ReadFile("problem.md")
	assert.Nil(t, err)
	t.Logf("\nProblem: %s \n[#tests=%d]", bytes, len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")
	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {

			actual := impl.canFinish(tt.n, tt.preReqs)

			var pOrF string
			if assert.Equal(t, tt.out, actual) {
				pOrF = solve_common_problems_using_go.Passed
			} else {
				pOrF = solve_common_problems_using_go.Failed
			}

			t.Logf("test # %d, %s input(s): [n=%v, pre-reqs=%v] and expected output=%v, %v", i+1, pOrF, tt.n, tt.preReqs, tt.out, tt.msg)
		}
	}
}
