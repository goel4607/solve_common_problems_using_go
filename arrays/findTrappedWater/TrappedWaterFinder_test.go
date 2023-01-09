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
				[]int{1, 3, 6, 9, 2},
				15,
			},
			{
				"Another positive test case",
				[]int{1, 3, 6, 9, 2},
				11,
			},
			{
				"A negative test case when no nums sum is same as target",
				[]int{1, 3, 6, 9, 2},
				25,
			},
			{
				"A negative test case when only one num is given",
				[]int{5},
				5,
			},
			{
				"A negative test case when no i/p is given",
				[]int{},
				25,
			},
		}

		for i, tt := range inputs {
			t.Logf("test # %d, input=%v, target=%v and expected output=%v, %v", i, tt.input, tt.expectedOpt, tt.info)
			actualOpt := impl.findTrappedWater(tt.input)
			assert.Equal(t, tt.expectedOpt, actualOpt)
		}
	}
}
