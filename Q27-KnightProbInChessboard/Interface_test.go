package minCostOfClimbingStairs

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
	n            int
	k            int
	r            int
	c            int
	expectedProb float64
}

func getTests() []InterfaceTest {
	return []InterfaceTest{
		{
			n:            3,
			k:            2,
			r:            0,
			c:            0,
			expectedProb: 0.06250,
		},
		{
			n:            1,
			k:            0,
			r:            0,
			c:            0,
			expectedProb: 1.00000,
		},
		{
			n:            3,
			k:            3,
			r:            0,
			c:            0,
			expectedProb: 0.015625,
		},
		{
			n:            10,
			k:            13,
			r:            5,
			c:            3,
			expectedProb: 0.11733772067236714,
		},
	}
}

func TestKnightProbInChessboard(t *testing.T) {
	// print the problem
	bytes, err := os.ReadFile("problem.md")
	assert.Nil(t, err)
	t.Logf("Problem: %s", bytes)

	impls = append(
		impls,
		&Soln1TopDown{},
		//&Soln1MemUsingArr{},
		//&Soln1{},
		//&Soln1TopDown{},
		//&Soln1TopDown{},
	)

	tests := getTests()
	t.Logf("[#tests=%d]", len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")
	for i, tt := range tests {
		t.Logf("\ttest %d\t%#v", i+1, tt)

		for _, impl := range impls {
			t.Logf("Using: %T", impl)

			actualMinCost := impl.computeProbabilityThatKnightIsStillOnCB(tt.n, tt.k, tt.r, tt.c)

			if assert.Equal(t, tt.expectedProb, actualMinCost) {
				t.Logf("\t%s\texpected output=%v, actual=%v", solve_common_problems_using_go.Passed, tt.expectedProb, actualMinCost)
			} else {
				t.Logf("\t%s\txpected output=%v, actual=%v", solve_common_problems_using_go.Failed, tt.expectedProb, actualMinCost)
			}
		}
	}
}
