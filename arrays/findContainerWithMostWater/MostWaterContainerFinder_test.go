package findContainerWithMostWater

import (
	"testing"
)
import "github.com/stretchr/testify/assert"

var (
	impls []MostWaterContainerFinder
)

func TestContainerWithMostWater(t *testing.T) {
	impls = append(impls, MostWaterContainerFinderS1BF{}, MostWaterContainerFinderS2Efficient{})

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		inputs := []struct {
			info        string
			input       []int
			expectedOpt int
		}{
			{
				"A positive test case with greatest value is among the area walls",
				[]int{7, 1, 2, 3, 9},
				28,
			},
			{
				"A positive test case with greatest value is among the area walls",
				[]int{1, 8, 6, 2, 9, 4},
				24,
			},
			{
				"Another positive test case with middle value greatest but it is not amongst the wall",
				[]int{1, 8, 6, 2, 9, 7},
				28,
			},
			{
				"Another positive test case when diff is only one",
				[]int{1, 9, 9, 9, 2},
				18,
			},
			{
				"A negative test case when only one num is given",
				[]int{5},
				0,
			},
			{
				"A negative test case when no i/p is given",
				[]int{},
				0,
			},
		}

		for i, tt := range inputs {
			t.Logf("test # %d, input=%v and expected output=%v, %v", i, tt.input, tt.expectedOpt, tt.info)
			actualOpt := impl.findMostWaterContainer(tt.input)
			assert.Equal(t, tt.expectedOpt, actualOpt)
		}
	}
}
