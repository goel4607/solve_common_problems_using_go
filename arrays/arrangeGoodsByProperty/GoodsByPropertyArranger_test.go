package arrangeGoodsByProperty

import (
	"github.com/solve_common_problems_using_go"
	"github.com/stretchr/testify/assert"
	"testing"
)

var problemStatement = `
Let’s say we have a warehouse, that contains following things with expiry dates:
- Carrot that will expire in 30 days,
- Potato for 45 days,
- Rice for 100 Days and
- Spinach can be used in 5 days.

So to deliver these items, we must take care their expiry time, and to avoid the expiration, we will be delivering
these goods from the lowest to the highest expiry time.

So the order should be:

Item        Expiry in days
--------------------------
Carrot      30
Spinach     5
Rice        100
Potato      45

Expected:
--------------------------
Spinach     5
Carrot      30
Potato      45
Rice        100
`

type ArrangeGoodsByPropertyTest struct {
	testSummary string
	data        []Item
	property    string
	expOpt      []Item
}

func testData() []ArrangeGoodsByPropertyTest {
	return []ArrangeGoodsByPropertyTest{
		{
			"normal test data with m and m in-between values",
			[]Item{
				{"Carrot", 30},
				{"Spinach", 5},
				{"Rice", 100},
				{"Potato", 45},
			},
			"",
			[]Item{
				{"Spinach", 5},
				{"Carrot", 30},
				{"Potato", 45},
				{"Rice", 100},
			},
		},
		{
			"normal test data with m and m in-between values",
			[]Item{
				{"Carrot", 30},
				{"Spinach", 5},
				{"Rice", 100},
				{"Potato", 45},
			},
			"name",
			[]Item{
				{"Carrot", 30},
				{"Potato", 45},
				{"Rice", 100},
				{"Spinach", 5},
			},
		},
	}
}

var (
	impls = make([]GoodsByPropertyArranger, 0)
)

func TestArrangeGoodsByPropertyImplementations(t *testing.T) {
	impls =
		append(impls,
			&GoodsByPropertyArrangerSolution1{},
		)
	tests := testData()

	t.Logf("%s\n[#tests=%d]", problemStatement, len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")
	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {

			actual := impl.ArrangeGoodsByProperty(tt.data, tt.property)
			var pOrF string
			if assert.Equal(t, tt.expOpt, actual) {
				pOrF = solve_common_problems_using_go.Passed
			} else {
				pOrF = solve_common_problems_using_go.Failed
			}

			t.Logf("test # %d, %s input[data=%v] and expected output=%v, %v", i+1, pOrF, tt.data, tt.expOpt, tt.testSummary)
		}
	}
}
