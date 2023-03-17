package interfaceDesignMonarchy

import (
	"github.com/solve_common_problems_using_go"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	impls []Monarchy
)

type MonarchyTest struct {
	birthOrder                [][2]string
	deathOrder                []string
	expectedOrderOfSuccession []string
}

func getTests() []MonarchyTest {
	return []MonarchyTest{
		{
			birthOrder: [][2]string{
				{"Jake", ""},
				{"Catherine", "Jake"},
				{"Jane", "Catherine"},
				{"Farah", "Jane"},
				{"Tom", "Jake"},
				{"Celine", "Jake"},
				{"Mark", "Catherine"},
				{"Peter", "Celine"},
			},
			deathOrder:                []string{"Jake", "Jane"},
			expectedOrderOfSuccession: []string{"Catherine", "Farah", "Mark", "Tom", "Celine", "Peter"},
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
		&Soln1{},
	)

	tests := getTests()
	t.Logf("[#tests=%d]", len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")
	for i, tt := range tests {
		t.Logf("\ttest %d\tmonarchy: \n\t\t\tbirthOrder=%+v, \n\t\t\tdeathOrder=%+v, \n\t\t\torderOfSuccession=%+v", i+1, tt.birthOrder, tt.deathOrder, tt.expectedOrderOfSuccession)

		for _, impl := range impls {
			t.Logf("Using: %T", impl)

			for _, b := range tt.birthOrder {
				err := impl.Birth(b[0], b[1])
				assert.Nil(t, err)
			}

			for _, d := range tt.deathOrder {
				err := impl.Death(d)
				assert.Nil(t, err)
			}

			succOrder := impl.GetOrderOfSuccession()

			if assert.Equal(t, tt.expectedOrderOfSuccession, succOrder) {
				t.Logf("\t%s\texpected output=%v, actual=%v", solve_common_problems_using_go.Passed, tt.expectedOrderOfSuccession, succOrder)
			} else {
				t.Logf("\t%s\txpected output=%v, actual=%v", solve_common_problems_using_go.Failed, tt.expectedOrderOfSuccession, succOrder)
			}
		}
	}
}
