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
	cost            []int
	expectedMinCost int
}

func getTests() []InterfaceTest {
	return []InterfaceTest{
		{
			cost:            []int{10, 15, 30},
			expectedMinCost: 15,
		},
		{
			cost:            []int{20, 15, 30, 5},
			expectedMinCost: 20,
		},
	}
}

func TestBSTValidator(t *testing.T) {
	// print the problem
	bytes, err := os.ReadFile("problem.md")
	assert.Nil(t, err)
	t.Logf("Problem: %s", bytes)

	impls = append(
		impls,
		//&Soln1WithoutMemoization{},
		//&Soln1MemUsingArr{},
		//&Soln1{},
		//&Soln1BtmUp{},
		&Soln1BtmUpWithMem{},
	)

	tests := getTests()
	t.Logf("[#tests=%d]", len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")
	for i, tt := range tests {
		t.Logf("\ttest %d\tinput: cost=%+v, \texpectedMinCost=%v", i+1, tt.cost, tt.expectedMinCost)

		for _, impl := range impls {
			t.Logf("Using: %T", impl)

			actualMinCost, numCalls := impl.findMinCostOfClimbingStairs(tt.cost)

			if assert.Equal(t, tt.expectedMinCost, actualMinCost) {
				t.Logf("\t%s\texpected output=%v, actual=%v, numCalls=%v", solve_common_problems_using_go.Passed, tt.expectedMinCost, actualMinCost, numCalls)
			} else {
				t.Logf("\t%s\txpected output=%v, actual=%v, numCalls=%v", solve_common_problems_using_go.Failed, tt.expectedMinCost, actualMinCost, numCalls)
			}
		}
	}
}
