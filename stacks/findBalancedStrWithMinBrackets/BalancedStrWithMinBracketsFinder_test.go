package findBalancedStrWithMinBrackets

import (
	"github.com/solve_common_problems_using_go"
	"github.com/stretchr/testify/assert"
	"testing"
)

type BalancedStrWithMinBracketsTest struct {
	testSummary string
	data        string
	expOpt      []string
}

func testData() []BalancedStrWithMinBracketsTest {
	return []BalancedStrWithMinBracketsTest{
		{
			"+ve test with ",
			"abc(ab(cd)ef",
			[]string{"abcab(cd)ef", "abc(abcd)ef"},
		},
		{
			"+ve test with ",
			"a(b)cd",
			[]string{"a(b)cd"},
		},
		{
			"+ve test with ",
			"(abc((cd)",
			[]string{"abc(cd)"},
		},
		{
			"+ve test with ",
			"abc)de(f)",
			[]string{"abcde(f)"},
		},
		{
			"+ve test with ",
			"))((",
			[]string{""},
		},
		{
			"+ve test with ",
			"a))b((c",
			[]string{"abc"},
		},
	}
}

var (
	impls []BalancedStrWithMinBracketsFinder
)

func TestMNReverserImplementations(t *testing.T) {
	impls =
		append(impls,
			//Solution1{},
			//Solution2{},
			Prac23Week11{},
		)
	tests := testData()
	t.Logf("Given a string only containing round brackets '(' and ')' and lowercase characters, remove the least"+
		"\namount of brackets such that the string is valid. [#tests=%d]", len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")
	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actual := impl.findBalancedStrWithMinBrackets(tt.data)

			var found bool
			for _, e := range tt.expOpt {
				if e == actual {
					found = true
					break
				}
			}
			var pOrF string
			if assert.True(t, found) {
				pOrF = solve_common_problems_using_go.Passed
			} else {
				pOrF = solve_common_problems_using_go.Failed
			}

			t.Logf("test # %d, %s input[data=%v] and expected output(any)=%v, %v", i+1, pOrF, tt.data, tt.expOpt, tt.testSummary)
		}
	}
}
