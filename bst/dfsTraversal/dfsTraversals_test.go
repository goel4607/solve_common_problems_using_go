package dfsTraversal

import (
	"github.com/solve_common_problems_using_go"
	"github.com/solve_common_problems_using_go/bst"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	impls []DfsTraversals
)

func TestCountHeight(t *testing.T) {

	tests := []struct {
		input                   []int
		expectedHeight          int
		expectedPreorderValues  []int
		expectedInorderValues   []int
		expectedPostorderValues []int
	}{
		{
			input:                   []int{1, 2, 3, 4, 5},
			expectedHeight:          5,
			expectedPreorderValues:  []int{1, 2, 3, 4, 5},
			expectedInorderValues:   []int{1, 2, 3, 4, 5},
			expectedPostorderValues: []int{5, 4, 3, 2, 1},
		},
		{
			input:                   []int{20, 10, 5, 6, 7, 8, 9},
			expectedHeight:          7,
			expectedPreorderValues:  []int{20, 10, 5, 6, 7, 8, 9},
			expectedInorderValues:   []int{5, 6, 7, 8, 9, 10, 20},
			expectedPostorderValues: []int{9, 8, 7, 6, 5, 10, 20},
		},
		{
			input:                   []int{20, 10, 25, 5, 12, 22, 30},
			expectedHeight:          3,
			expectedPreorderValues:  []int{20, 10, 5, 12, 25, 22, 30},
			expectedInorderValues:   []int{5, 10, 12, 20, 22, 25, 30},
			expectedPostorderValues: []int{5, 12, 10, 22, 30, 25, 20},
		},
	}

	impls = append(impls, &S1DfsTraversals{})

	for _, impl := range impls {
		for i, tt := range tests {
			t.Logf("\ttest %d\tinput=%+v, expectedValues: Preorder=%v, Inorder=%v, Postorder=%v", i+1, tt.input, tt.expectedPreorderValues, tt.expectedInorderValues, tt.expectedPostorderValues)

			bst := &bst.Node{Data: tt.input[0]}
			for i, d := range tt.input {
				if i != 0 {
					bst.Insert(d)
				}
			}

			actual := impl.Preorder(bst)
			if assert.Equal(t, tt.expectedPreorderValues, actual) {
				t.Logf("\t%s\texpected Preorder Values=%v, actual=%v", solve_common_problems_using_go.Passed, tt.expectedPreorderValues, actual)
			} else {
				t.Logf("\t%s\texpected Preorder Values=%v, actual=%v", solve_common_problems_using_go.Failed, tt.expectedPreorderValues, actual)
			}

			actual = impl.Inorder(bst)
			if assert.Equal(t, tt.expectedInorderValues, actual) {
				t.Logf("\t%s\texpected Inorder Values=%v, actual=%v", solve_common_problems_using_go.Passed, tt.expectedInorderValues, actual)
			} else {
				t.Logf("\t%s\texpected Inorder Values=%v, actual=%v", solve_common_problems_using_go.Failed, tt.expectedInorderValues, actual)
			}

			actual = impl.Postorder(bst)
			if assert.Equal(t, tt.expectedPostorderValues, actual) {
				t.Logf("\t%s\texpected Postorder Values=%v, actual=%v", solve_common_problems_using_go.Passed, tt.expectedPostorderValues, actual)
			} else {
				t.Logf("\t%s\texpected Postorder Values=%v, actual=%v", solve_common_problems_using_go.Failed, tt.expectedPostorderValues, actual)
			}
		}
	}
}
