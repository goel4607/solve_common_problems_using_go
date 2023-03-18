package tries

type Trie interface {
	Insert(s string)
	Search(s string) bool
	StartsWith(prefix string) bool
}
