package course_scheduler

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
	msg     string
	n       int
	preReqs []CourseDependency
	out     bool
}

func getTests() []InterfaceTest {
	return []InterfaceTest{
		{
			msg: "+ve, course dependency with no cycle.",
			preReqs: []CourseDependency{
				0: {1, 0},
				1: {2, 1},
				2: {2, 5},
				3: {0, 3},
				4: {3, 5},
				5: {4, 3},
				6: {4, 5},
			},
			out: true,
		},
		{
			msg: "+ve, course dependency with one cycle.",
			preReqs: []CourseDependency{
				0: {1, 0},
				1: {2, 1},
				2: {5, 2},
				3: {0, 3},
				4: {3, 5},
				5: {4, 3},
				6: {4, 5},
			},
			out: true,
		},
		{
			msg: "-ve, course dependency with one cycle.",
			preReqs: []CourseDependency{
				0: {1, 0},
				1: {2, 1},
				2: {2, 5},
				3: {0, 3},
				4: {3, 5},
				5: {4, 3},
				6: {5, 4},
			},
			out: true,
		},
		{
			msg: "-ve, course dependency with two cycles.",
			preReqs: []CourseDependency{
				0: {1, 0},
				1: {2, 1},
				2: {5, 2},
				3: {0, 3},
				4: {3, 5},
				5: {4, 3},
				6: {5, 4},
			},
			out: true,
		},
	}
}

func TestInterface(t *testing.T) {
	impls = append(
		impls,
		//Soln1{},
	)

	tests := getTests()
	bytes, err := os.ReadFile("problem.md")
	assert.Nil(t, err)
	t.Logf("\nProblem: %s \n[#tests=%d]", bytes, len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")
	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {

			impl.isItPossibleToFinishAllCourses(tt.n, tt.preReqs)

			var pOrF string
			if assert.Equal(t, tt.out, tt.preReqs) {
				pOrF = solve_common_problems_using_go.Passed
			} else {
				pOrF = solve_common_problems_using_go.Failed
			}

			t.Logf("test # %d, %s input=%v and expected output=%v, %v", i+1, pOrF, tt.preReqs, tt.out, tt.msg)
		}
	}
}
