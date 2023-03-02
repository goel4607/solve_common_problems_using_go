package findAllOrangesRottenTime

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
	in  [][]int
	out int
}

func getTests() []InterfaceTest {
	return []InterfaceTest{
		{
			msg: "5 by 4 array",
			in: [][]int{
				{2, 1, 1, 0, 0},
				{1, 1, 1, 0, 0},
				{0, 1, 1, 1, 1},
				{0, 1, 0, 0, 1},
			},
			out: 7,
		},
		{
			msg: "5 by 4 array",
			in: [][]int{
				{1, 1, 0, 0, 0},
				{2, 1, 0, 0, 0},
				{0, 0, 0, 1, 2},
				{0, 1, 0, 0, 1},
			},
			out: -1,
		},
		{
			msg: "empty array with no elements",
			in:  [][]int{},
			out: 0,
		},
		{
			msg: "empty array",
			in: [][]int{
				{},
				{},
			},
			out: 0,
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
	t.Logf("Problem: \n%s \n[#tests=%d]", bytes, len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")
	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {

			a := impl.findAllOrangesRottenTime(tt.in)

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
