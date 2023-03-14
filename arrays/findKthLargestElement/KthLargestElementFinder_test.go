package findKthLargestElement

import (
	"github.com/solve_common_problems_using_go"
	"github.com/stretchr/testify/assert"
	"testing"
)

type LargestElementFinderTest struct {
	info        string
	input       []int
	pos         int
	expectedOpt int
}

func testData() []LargestElementFinderTest {
	return []LargestElementFinderTest{
		{
			"+v test with no duplicate values",
			[]int{4, 5, 1, 2, 8, 7},
			3,
			5,
		},
		{
			"+v test with no duplicate values",
			[]int{5, 3, 1, 6, 4, 2},
			2,
			5,
		},
		{
			"+v test with duplicate values",
			[]int{2, 3, 1, 2, 4, 2},
			4,
			2,
		},
		{
			"+v test with no duplicate values but edge case",
			[]int{4, 5, 1, 2, 8, 7},
			6,
			1,
		},
		{
			"+v test with duplicate values",
			[]int{4, 7, 5, 1, 2, 7, 8, 7},
			4,
			7,
		},
		{
			"+v test with duplicate values edge case",
			[]int{4, 7, 5, 1, 2, 7, 7},
			1,
			7,
		},
		{
			"+v test with duplicate values edge case",
			[]int{4, 7, 5, 1, 2, 7, 7},
			7,
			1,
		},
	}
}

var (
	impls []LargestElementFinder
)

func TestLargestElementFinderAtPos(t *testing.T) {
	impls = append(
		impls,
		//S1BF{},
		//S2UsingQSelect{},
		//&Prac1{},
		//&Prac2UsingQS{},
		//&Prac23Week11UsingPQ{},
		&Prac23Week11UsingQSelect{},
	)

	tests := testData()
	t.Logf("Given an unsorted array, return the kth largest element. It is the kth largest element in sorted order, not the kth distinct element [#tests=%d]", len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {

			actualOpt := impl.findKthLargestElement(tt.input, tt.pos)
			var pOrF string
			if assert.Equal(t, tt.expectedOpt, actualOpt) {
				pOrF = solve_common_problems_using_go.Passed
			} else {
				pOrF = solve_common_problems_using_go.Failed
			}

			t.Logf("test # %d, %s input=[%v,k=%v] and expected output=%v, %v", i+1, pOrF, tt.input, tt.pos, tt.expectedOpt, tt.info)
		}
	}
}
