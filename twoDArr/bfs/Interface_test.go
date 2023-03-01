package bfs

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
	r   int
	c   int
	out []int
}

func getTests() []InterfaceTest {
	return []InterfaceTest{
		{
			msg: "5 by 5 array, with inc elements on each row",
			in: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
			},
			r:   2,
			c:   2,
			out: []int{13, 8, 14, 18, 12, 3, 9, 7, 15, 19, 17, 11, 4, 2, 10, 6, 20, 16, 5, 1},
		},
		{
			msg: "5 by 5 array, with inc elements on each row",
			in: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25},
			},
			r:   2,
			c:   2,
			out: []int{13, 8, 14, 18, 12, 3, 9, 7, 15, 19, 23, 17, 11, 4, 2, 10, 6, 20, 24, 22, 16, 5, 1, 25, 21},
		},
		{
			msg: "5 by 5 array, with inc elements on each row",
			in: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25},
			},
			r:   0,
			c:   0,
			out: []int{1, 2, 6, 3, 7, 11, 4, 8, 12, 16, 5, 9, 13, 17, 21, 10, 14, 18, 22, 15, 19, 23, 20, 24, 25},
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
	t.Logf("Problem: %s [#tests=%d]", bytes, len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")
	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {

			a := impl.BFS(tt.in, tt.r, tt.c)

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
