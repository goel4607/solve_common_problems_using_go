package findTypedOutStrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	impls = make([]TypedOutStringsEqualsFinder, 0)
)

func TestTypedOutStringsEqualFinder(t *testing.T) {
	impls = append(impls, TypedOutStringsEqualsFinderS1Bf{})
	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		inputs := []struct {
			msg         string
			s           string
			t           string
			expectedOpt bool
		}{
			{"+ve test with both strings have multiple continuous hashes", "bcd###", "abcd###", false},
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

		for i, tt := range inputs {
			t.Logf("test # %d, input=%v, %v, and expected output=%v, %v", i, tt.s, tt.t, tt.expectedOpt, tt.msg)
			actualOpt := impl.findIfTypedOutStringsAreEqual(tt.s, tt.t)
			assert.Equal(t, tt.expectedOpt, actualOpt)
		}
	}
}
