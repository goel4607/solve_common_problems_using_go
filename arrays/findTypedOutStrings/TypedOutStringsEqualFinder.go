package findTypedOutStrings

const hash = '#'

type TypedOutStringsEqualsFinder interface {
	findIfTypedOutStringsAreEqual(s, t string) bool
}
