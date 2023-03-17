package binSearch

import (
	"github.com/solve_common_problems_using_go"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	impls = make([]Interface, 0)
)

type InterfaceTest struct {
	msg string
	in  []int
	val int
	out int
}

func getTests() []InterfaceTest {
	return []InterfaceTest{
		{
			msg: "a normal sorted ints",
			in:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			val: 6,
			out: 5,
		},
		{
			msg: "a normal sorted ints",
			in:  []int{1, 2, 3, 4, 5, 5, 5, 5, 6, 7, 8},
			val: 6,
			out: 8,
		},
		{
			msg: "a normal sorted ints",
			in:  []int{1, 3, 5, 5, 5, 5, 6, 7, 8},
			val: 3,
			out: 1,
		},
	}
}

func TestInterface(t *testing.T) {
	impls = append(
		impls,
		Soln1{},
	)

	tests := getTests()
	bytes, err := os.ReadFile("problem.md")
	assert.Nil(t, err)
	t.Logf("Problem: %s \n[#tests=%d]", bytes, len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")
	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {

			a := impl.searchIndexInSortedArray(tt.in, tt.val)

			var pOrF string
			if assert.Equal(t, tt.out, a) {
				pOrF = solve_common_problems_using_go.Passed
			} else {
				pOrF = solve_common_problems_using_go.Failed
			}

			t.Logf("test # %d, %s input=%v and expected output=%v, %v", i+1, pOrF, tt.in, tt.out, tt.msg)
		}
	}
}
