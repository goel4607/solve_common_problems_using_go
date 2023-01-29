package useStackAsQueue

import (
	"github.com/solve_common_problems_using_go"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	impls []Queue
)

type QueueTest struct {
	testSummary string
	data        []int
	expOpt      []int
}

func testData() []QueueTest {
	return []QueueTest{
		{
			"+ve test with ",
			[]int{123456},
			[]int{123456},
		},
		{
			"+ve test with ",
			[]int{1},
			[]int{1},
		},
		{
			"+ve test with ",
			[]int{},
			[]int{},
		},
	}
}

func TestQuestImplementations(t *testing.T) {
	impls =
		append(impls,
			&Solution1{},
		)

	tests := testData()
	t.Logf("Simple Q implementation [#tests=%d]", len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")
	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			t.Logf("test # %d, queue functinality: input[data=%v] and expected output(any)=%v, %v", i+1, tt.data, tt.expOpt, tt.testSummary)
			for _, v := range tt.data {
				impl.enqueue(v)

				pOrF := solve_common_problems_using_go.Failed
				a, aExist := impl.peek()
				if assert.True(t, aExist) && assert.Equal(t, a, v) {
					pOrF = solve_common_problems_using_go.Passed
				}
				t.Logf("\t\t%v checking: enqueue() and peek(): input[data=%v] and expected output(any)=%v", pOrF, v, v)

				pOrF = solve_common_problems_using_go.Failed
				if assert.False(t, impl.empty()) {
					pOrF = solve_common_problems_using_go.Passed
				}
				t.Logf("\t\t%v checking: empty(): input[data=%v] and expected output(any)=%v", pOrF, v, v)

				pOrF = solve_common_problems_using_go.Failed
				a, aExist = impl.dequeue()
				if assert.True(t, aExist) && assert.Equal(t, a, v) {
					pOrF = solve_common_problems_using_go.Passed
				}
				t.Logf("\t\t%v checking: dequeue() and empty(): input[data=%v] and expected output(any)=%v", pOrF, v, v)
			}

			for _, v := range tt.data {
				impl.enqueue(v)
			}

			var actual = make([]int, 0)
			for i := 0; i < len(tt.data); i++ {
				if v, ok := impl.dequeue(); ok {
					actual = append(actual, v)
				}
			}

			pOrF := solve_common_problems_using_go.Failed
			if assert.Equal(t, actual, tt.expOpt) {
				pOrF = solve_common_problems_using_go.Passed
			}

			t.Logf("\t\t%v checking: enqueue() and dequeue() -> input[data=%v] and expected output(any)=%v", pOrF, tt.data, tt.expOpt)

			pOrF = solve_common_problems_using_go.Failed
			_, aExist := impl.peek()
			if assert.False(t, aExist) {
				pOrF = solve_common_problems_using_go.Passed
			}
			t.Logf("\t\t%v checking: peek() -> input[data=%v] and expected output(any)=%v", pOrF, nil, nil)

			pOrF = solve_common_problems_using_go.Failed
			_, aExist = impl.dequeue()
			if assert.False(t, aExist) {
				pOrF = solve_common_problems_using_go.Passed
			}
			t.Logf("\t\t%v checking: dequeue() -> input[data=%v] and expected output(any)=%v", pOrF, nil, nil)

			pOrF = solve_common_problems_using_go.Failed
			if assert.True(t, impl.empty()) {
				pOrF = solve_common_problems_using_go.Passed
			}
			t.Logf("\t\t%v checking: empty() -> input[data=%v] and expected output(any)=%v", pOrF, nil, true)
		}
	}
}
