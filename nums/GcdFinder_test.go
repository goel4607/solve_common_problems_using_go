package nums

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var impls []GcdFinder

type GcdFinderTestData struct {
	info        string
	a           int
	b           int
	expectedOpt int
}

func testsToRun() []GcdFinderTestData {
	return []GcdFinderTestData{
		{"main test", 78, 66, 6},
		{"no common gcd", 78, 67, 1},
	}
}

func TestGcdFinder(t *testing.T) {
	impls = append(
		impls,
		GcdFinderSoln1{},
	)

	assert.NotEmptyf(t, impls, "no implemantations present!!")

	tests := testsToRun()

	for i := 0; i < len(impls); i++ {
		for j, tt := range tests {
			t.Logf("test # %d, input values=[a=%v, b=%v, expected output=%v, msg=%v", j, tt.a, tt.b, tt.expectedOpt, tt.info)
			actualOpt := impls[i].findGcd(tt.a, tt.b)
			assert.Equal(t, tt.expectedOpt, actualOpt, fmt.Sprintf("expect=%v, actual=%v", tt.expectedOpt, actualOpt))
		}
	}
}
