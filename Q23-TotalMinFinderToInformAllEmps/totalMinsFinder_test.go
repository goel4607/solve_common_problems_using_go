package Q23_TotalMinFinderToInformAllEmps

import (
	"github.com/solve_common_problems_using_go"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	impls = make([]TotalMinutesFinderToInformAllEmps, 0)
)

type InterfaceTest struct {
	msg               string
	headID            int
	managers          []int
	informTime        []int
	expectedTotalTime int
}

func getTests() []InterfaceTest {
	return []InterfaceTest{
		{
			msg:    "head has 3 employees and each of those has 1 or 2 employees",
			headID: 4,
			managers: []int{
				0: 2,
				1: 2,
				2: 4,
				3: 6,
				4: -1,
				5: 4,
				6: 4,
				7: 5,
			},
			informTime: []int{
				0: 0,
				1: 0,
				2: 4,
				3: 0,
				4: 7,
				5: 3,
				6: 6,
				7: 0,
			},
			expectedTotalTime: 13,
		},
		{
			msg:    "here each manager has only one employee which means that each level of the tree has only one node so there is max calculation i.e. its an addition of times only.",
			headID: 6,
			managers: []int{
				0: 1,
				1: 2,
				2: 3,
				3: 4,
				4: 5,
				5: 6,
				6: -1,
			},
			informTime: []int{
				0: 0,
				1: 6,
				2: 5,
				3: 4,
				4: 3,
				5: 2,
				6: 1,
			},
			expectedTotalTime: 21,
		},
		{
			msg:    "no managers",
			headID: 0,
			managers: []int{
				0: -1,
			},
			informTime: []int{
				0: 0,
			},
			expectedTotalTime: 0,
		},
	}
}

func TestInterface(t *testing.T) {
	impls = append(
		impls,
		//Soln1UsingBfs{},
		//Prac23Week11UsingBfs{},
		//Prac23Week11UsingDFS{},
		Prac23AprDFS{},
	)

	tests := getTests()
	bytes, err := os.ReadFile("problem.md")
	assert.Nil(t, err)
	t.Logf("Problem: %s [#tests=%d]", bytes, len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")
	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {

			aTotalTime := impl.FindTotalMinutesToInformAllEmps(tt.headID, tt.managers, tt.informTime)

			var pOrF string
			if assert.Equal(t, tt.expectedTotalTime, aTotalTime) {
				pOrF = solve_common_problems_using_go.Passed
			} else {
				pOrF = solve_common_problems_using_go.Failed
			}

			t.Logf("test # %d, %s input=(managers=%v, informTime=%v) and expected output=%v, %v", i+1, pOrF, tt.managers, tt.informTime, tt.expectedTotalTime, tt.msg)
		}
	}
}
