package sodokuSolver

import (
	"github.com/solve_common_problems_using_go"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	impls []Interface
)

type InterfaceTest struct {
	puzzle [][]int
}

func getTests() []InterfaceTest {
	return []InterfaceTest{
		{
			puzzle: [][]int{
				0: {5, 3, 0, 0, 7, 0, 0, 0, 0},
				1: {6, 0, 0, 1, 9, 5, 0, 0, 0},
				2: {0, 9, 8, 0, 0, 0, 0, 6, 0},
				3: {8, 0, 0, 0, 6, 0, 0, 0, 3},
				4: {4, 0, 0, 8, 0, 3, 0, 0, 1},
				5: {7, 0, 0, 0, 2, 0, 0, 0, 6},
				6: {0, 6, 0, 0, 0, 0, 2, 8, 0},
				7: {0, 0, 0, 4, 1, 9, 0, 0, 5},
				8: {0, 0, 0, 0, 8, 0, 0, 7, 9},
			},
		},
	}
}

func TestMinCostOfClimbingStairs(t *testing.T) {
	// print the problem
	bytes, err := os.ReadFile("problem.md")
	assert.Nil(t, err)
	t.Logf("Problem: %s", bytes)

	impls = append(
		impls,
		&Soln1{},
	)

	tests := getTests()
	t.Logf("[#tests=%d]", len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")
	for i, tt := range tests {
		t.Logf("\ttest %d\tinput: cost=%+v", i+1, tt.puzzle)

		for _, impl := range impls {
			t.Logf("Using: %T", impl)

			impl.SodukuSolver(tt.puzzle)

			if verify(t, tt.puzzle) {
				t.Logf("\t%s\tactual=%v", solve_common_problems_using_go.Passed, tt.puzzle)
			} else {
				t.Logf("\t%s\tactual=%v", solve_common_problems_using_go.Failed, tt.puzzle)
			}
		}
	}
}

func verify(t *testing.T, soln [][]int) bool {
	aOccupation := [9][10]bool{}
	eOccupation := [9][10]bool{
		{false, true, true, true, true, true, true, true, true, true},
		{false, true, true, true, true, true, true, true, true, true},
		{false, true, true, true, true, true, true, true, true, true},
		{false, true, true, true, true, true, true, true, true, true},
		{false, true, true, true, true, true, true, true, true, true},
		{false, true, true, true, true, true, true, true, true, true},
		{false, true, true, true, true, true, true, true, true, true},
		{false, true, true, true, true, true, true, true, true, true},
		{false, true, true, true, true, true, true, true, true, true},
	}

	for i, row := range soln {
		for _, v := range row {
			aOccupation[i][v] = true
		}
	}

	return assert.Equal(t, eOccupation, aOccupation)
}
