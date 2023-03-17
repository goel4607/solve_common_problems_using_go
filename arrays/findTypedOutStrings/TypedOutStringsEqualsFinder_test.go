package findTypedOutStrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	impls = make([]TypedOutStringsEqualsFinder, 0)
)

type TypedOutStringTest struct {
	msg         string
	s           string
	t           string
	expectedOpt bool
}

func getTests() []TypedOutStringTest {
	return []TypedOutStringTest{
		{"+ve test with strings differing in hashes", "a##c", "#a#c", true},
		{"-ve test with both strings have multiple continuous hashes", "bcd###", "abcd###", false},
		{"+ve test with both strings have same number of hashes", "ab#z", "az#z", true},
		{"-ve test", "abc#d", "acc#c", false},
		{"+ve test with multiple hashes in one and single hash in other string", "x#y#z#", "a#", true},
		{"+ve test with multiple hashes such their not sufficient chars to remove", "a###b", "#b", true},
		{"-ve test with uppercase/lowercase diff", "Ab#z", "ab#z", false},
		{"+ve test with tricky input", "bxj##tw", "bxo#j##tw", true},
		{"-ve test with tricky input", "bxj##tw", "bxj###tw", false},
		{"+ve test with tricky input", "y#fo##f", "y#f#o##f", true},
		{"+ve test with tricky input", "nzp#o#g", "b#nzp#o#g", true},
		{"-ve test with tricky input", "bbbextm", "bbb#extm", false},
	}
}

func TestTypedOutStringsEqualFinder(t *testing.T) {
	impls = append(
		impls,
		//TypedOutStringsEqualsFinderS1Bf{},
		//TypedOutStringsEqualsFindersS2Efficient{},
		//TypedOutStringsEqualsFindersS3EfficientSimplified{},
		//TypedOutStringsEqualFinderS4Prac1{},
		//TypedOutStringsEqualsFinderS4Prac2{},
		Prac23Week11{},
	)

	tests := getTests()
	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			t.Logf("test # %d, input=%v, %v, and expected output=%v, %v", i, tt.s, tt.t, tt.expectedOpt, tt.msg)
			actualOpt := impl.findIfTypedOutStringsAreEqual(tt.s, tt.t)
			assert.Equal(t, tt.expectedOpt, actualOpt)
		}
	}
}
