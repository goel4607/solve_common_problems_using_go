package Q30_ImplementPrefixTrie

type Trie interface {
	Insert(s string)
	Search(s string) bool
	StartsWith(prefix string) bool
}
