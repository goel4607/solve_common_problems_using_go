package Q4_FindTypedOutStrings

const hash = '#'

type TypedOutStringsEqualsFinder interface {
	findIfTypedOutStringsAreEqual(s, t string) bool
}
