package findtwosums

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"

var (
	impls []TwoSumFinder
)

func TestFindTwoSumsPositive(t *testing.T) {
	impls = append(impls, TwoSumFinderSolution1BF{}, TwoSumFinderS2Efficient{})

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		inputs := []struct {
			info        string
			vals        []int
			target      int
			expectedOpt []int
		}{
			{
				"A positive test case",
				[]int{1, 3, 6, 9, 2},
				15,
				[]int{2, 3},
			},
			{
				"Another positive test case",
				[]int{1, 3, 6, 9, 2},
				11,
				[]int{3, 4},
			},
			{
				"A negative test case when no nums sum is same as target",
				[]int{1, 3, 6, 9, 2},
				25,
				[]int{},
			},
			{
				"A negative test case when only one num is given",
				[]int{5},
				5,
				[]int{},
			},
			{
				"A negative test case when no i/p is given",
				[]int{},
				25,
				[]int{},
			},
		}

		for i, tt := range inputs {
			t.Logf("test # %d, input values=%v, target=%v and expected output=%v, %v", i, tt.vals, tt.target, tt.expectedOpt, tt.info)
			actualOpt := impl.findTwoSum(tt.vals, tt.target)
			assert.Equal(t, tt.expectedOpt, actualOpt, fmt.Sprintf("target=%v should be present", tt.target))
		}
	}
}
