package findTrappedWater

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	impls = []TrappedWaterFinder{}
)

func TestFindTrappedWater(t *testing.T) {
	//impls = {}
	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		inputs := []struct {
			info        string
			input       []int
			expectedOpt int
		}{
			{
				"A positive test case",
				[]int{0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2},
				8,
			},
			{
				"Another positive test case",
				[]int{0, 1, 0, 2, 1, 0, 3, 0, 0, 1, 2},
				9,
			},
			{
				"A negative test case when no value is captured as no trapped water",
				[]int{3, 4, 3},
				0,
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
			t.Logf("test # %d, input=%v, target=%v and expected output=%v, %v", i, tt.input, tt.expectedOpt, tt.info)
			actualOpt := impl.findTrappedWater(tt.input)
			assert.Equal(t, tt.expectedOpt, actualOpt)
		}
	}
}
