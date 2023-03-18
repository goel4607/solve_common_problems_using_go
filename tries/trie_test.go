package tries

import (
	"github.com/solve_common_problems_using_go"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	impls []Trie
)

type searchTest struct {
	s                string
	expectedPresence bool
}
type TrieTest struct {
	inserts         []string
	searchTests     []searchTest
	startsWithTests []searchTest
}

func getTests() []TrieTest {
	return []TrieTest{
		{
			inserts: []string{
				"apple",
				"dog",
			},
			searchTests: []searchTest{
				{"apple", true},
				{"dog", true},
				{"dogs", false},
				{"app", false},
			},
			startsWithTests: []searchTest{
				{"apple", true},
				{"dog", true},
				{"app", true},
			},
		},
		{
			inserts: []string{
				"apple",
				"dog",
				"app",
			},
			searchTests: []searchTest{
				{"apple", true},
				{"dog", true},
				{"apps", false},
				{"apples", false},
				{"dogs", false},
				{"app", true}, //new one
			},
			startsWithTests: []searchTest{
				{"a", true},
				{"ap", true},
				{"app", true},
				{"apps", false},
				{"appl", true},
				{"apple", true},
				{"apples", false},
				{"dog", true},
				{"app", true},
			},
		},
	}
}

func TestTrie(t *testing.T) {
	// print the problem
	bytes, err := os.ReadFile("problem.md")
	assert.Nil(t, err)
	t.Logf("Problem: %s", bytes)

	impls = append(
		impls,
	)

	tests := getTests()
	t.Logf("[#tests=%d]", len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")
	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			t.Logf("\ttest %d \tinput=[\n\t\t\tinserts=%+v, \n\t\t\tsearches=%+v, \n\t\t\tsearchPrefixes=%+v]", i+1, tt.inserts, tt.searchTests, tt.startsWithTests)

			for _, s := range tt.inserts {
				impl.Insert(s)
			}

			for _, st := range tt.searchTests {
				actual := impl.Search(st.s)
				if assert.Equal(t, st.expectedPresence, actual) {
					t.Logf("\t%s\texpected output=%v, actual=%v", solve_common_problems_using_go.Passed, st.expectedPresence, actual)
				} else {
					t.Logf("\t%s\txpected output=%v, actual=%v", solve_common_problems_using_go.Failed, st.expectedPresence, actual)
				}
			}
		}
	}
}
