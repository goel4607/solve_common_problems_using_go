package bstValidation

import (
	"github.com/solve_common_problems_using_go"
	"github.com/solve_common_problems_using_go/bst"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	impls []BSTValidator
)

type BSTValidatorTest struct {
	in      []int
	isValid bool
}

func getTests() []BSTValidatorTest {
	return []BSTValidatorTest{
		{
			in:      []int{1, 2, 4, 3, 5},
			isValid: false,
		},
		{
			in:      []int{20, 10, 5, 6, 7, 4, 9},
			isValid: false,
		},
		{
			in:      []int{20, 10, 25, 5, 12, 22, 30},
			isValid: true,
		},
		{
			in:      []int{20, 10, 25, 5, 12, 22, 30},
			isValid: true,
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
		//&Solution1BSTValidator{},
		Prac23Week11{},
	)

	tests := getTests()
	t.Logf("[#tests=%d]", len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")
	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			t.Logf("\ttest %d\tinput=%+v, expected isValid=%v", i+1, tt.in, tt.isValid)

			head := bst.NewWithValuesInSpecifiedOrder(tt.in)

			actual := impl.Validate(head)
			if assert.Equal(t, tt.isValid, actual) {
				t.Logf("\t%s\texpected output=%v, actual=%v", solve_common_problems_using_go.Passed, tt.isValid, actual)
			} else {
				t.Logf("\t%s\txpected output=%v, actual=%v", solve_common_problems_using_go.Failed, tt.isValid, actual)
			}
		}
	}
}
