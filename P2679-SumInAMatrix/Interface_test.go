package P679_SumInAMatrix

import (
	"fmt"
	"github.com/solve_common_problems_using_go"
	"github.com/stretchr/testify/assert"
	"math"
	"os"
	"testing"
)

var (
	impls []Interface
)

type InterfaceTest struct {
	nums        [][]int
	expectedSum int
}

func getTests() []InterfaceTest {
	return []InterfaceTest{
		{
			nums: [][]int{
				{7, 2, 1},
				{6, 4, 2},
				{6, 5, 3},
				{3, 2, 1},
			},
			expectedSum: 15,
		},
		{
			nums: [][]int{
				{1},
			},
			expectedSum: 1,
		},
		{
			nums: [][]int{
				{1, 8, 16, 15, 12, 9, 15, 11, 18, 6, 16, 4, 9, 4},
				{3, 19, 8, 17, 19, 4, 9, 3, 2, 10, 15, 17, 3, 11},
				{13, 10, 19, 20, 6, 17, 15, 14, 16, 8, 1, 17, 0, 2},
				{12, 20, 0, 19, 15, 10, 7, 10, 2, 6, 18, 7, 7, 4},
				{17, 14, 2, 2, 10, 16, 15, 3, 9, 17, 9, 3, 17, 10},
				{17, 6, 19, 17, 18, 9, 14, 2, 19, 12, 10, 18, 7, 9},
				{5, 6, 5, 1, 19, 8, 15, 2, 2, 4, 4, 1, 2, 17},
				{12, 16, 8, 16, 7, 6, 18, 13, 18, 8, 14, 15, 20, 11},
				{2, 10, 19, 3, 15, 18, 20, 10, 6, 7, 0, 8, 3, 7},
				{11, 5, 10, 13, 1, 3, 4, 7, 1, 18, 20, 17, 19, 2},
				{0, 3, 20, 6, 19, 18, 3, 12, 2, 11, 3, 1, 19, 0},
				{6, 5, 3, 15, 6, 1, 0, 17, 13, 19, 3, 8, 2, 7},
				{2, 20, 9, 11, 13, 5, 1, 16, 14, 1, 19, 3, 12, 6},
			},
			expectedSum: 190,
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
		Soln1UsingPQ{},
	)

	tests := getTests()
	t.Logf("[#tests=%d]", len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")
	for i, tt := range tests {
		t.Logf("\ttest %d\t%#v", i+1, tt)

		for _, impl := range impls {
			t.Logf("Using: %T", impl)

			actual := impl.matrixSum(tt.nums)

			if assert.Equal(t, tt.expectedSum, actual) {
				t.Logf("\t%s\texpected output=%v, actual=%v", solve_common_problems_using_go.Passed, tt.expectedSum, actual)
			} else {
				t.Logf("\t%s\txpected output=%v, actual=%v", solve_common_problems_using_go.Failed, tt.expectedSum, actual)
			}
		}
	}
}

func TestAny(t *testing.T) {
	v := int(math.Pow(2, 4))
	fmt.Println(v)
	vals := []int{12, 18}
	bOr := vals[0]
	for i := 1; i < len(vals); i++ {
		bOr = bOr | vals[i]
	}

	fmt.Println(bOr)
}
