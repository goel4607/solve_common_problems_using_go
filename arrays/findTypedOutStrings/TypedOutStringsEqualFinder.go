package findTypedOutStrings

type TypedOutStringsEqualsFinder interface {
	findIfTypedOutStringsAreEqual(s, t string) bool
}
