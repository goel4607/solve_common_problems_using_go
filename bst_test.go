package solve_common_problems_using_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	impl BSTNode
)

func TestCountHeight(t *testing.T) {

	tests := []struct {
		in                []int
		expectedSortedOpt []int
	}{
		{
			in:                []int{1, 2, 3, 4, 5},
			expectedSortedOpt: []int{1, 2, 3, 4, 5},
		},
		{
			in:                []int{20, 10, 5, 6, 7, 8, 9},
			expectedSortedOpt: []int{5, 6, 7, 8, 9, 10, 20},
		},
		{
			in:                []int{20, 10, 25, 5, 12, 22, 30},
			expectedSortedOpt: []int{5, 10, 12, 20, 22, 25, 30},
		},
	}

	for i, tt := range tests {
		t.Logf("\ttest %d\tinput=%+v, expected sorted values=%v", i+1, tt.in, tt.expectedSortedOpt)

		bst := &BSTNode{Data: tt.in[0]}
		for j, d := range tt.in {
			if j != 0 {
				bst.Insert(d)
			}
		}

		actual := bst.SortedList()
		if assert.Equal(t, tt.expectedSortedOpt, actual) {
			t.Logf("\t%s\texpected sorted output=%v, actual=%v", Passed, tt.expectedSortedOpt, actual)
		} else {
			t.Logf("\t%s\txpected sorted output=%v, actual=%v", Failed, tt.expectedSortedOpt, actual)
		}
	}
}
