package Q30_ImplementPrefixTrie

import "strings"

const (
	NumAlphabets  = 26
	StartingIndex = rune('a')
)

type TrieNode struct {
	c        rune
	isEnd    bool
	children [NumAlphabets]*TrieNode
}

func (n *TrieNode) getRelativePosition(c rune) int {
	return int(c - StartingIndex)
}

func (n *TrieNode) addChar(c rune, isEnding bool) *TrieNode {
	arrIdx := n.getRelativePosition(c)
	newChar := &TrieNode{
		c:     c,
		isEnd: isEnding,
	}

	n.children[arrIdx] = newChar

	return newChar
}

func (n *TrieNode) addCharIfNotPresent(c rune, isEnding bool) *TrieNode {
	arrIdx := n.getRelativePosition(c)

	if n.children[arrIdx] == nil {
		newChar := &TrieNode{
			c:     c,
			isEnd: isEnding,
		}

		n.children[arrIdx] = newChar
		return newChar
	}

	return n.children[arrIdx]
}

func (n *TrieNode) Insert(s string) {
	s = strings.ToLower(s)

	nextNode := n

	lastIdx := len(s) - 1
	for i, c := range s {
		var isEnd bool
		if i == lastIdx {
			isEnd = true
		}

		nextNode = nextNode.addCharIfNotPresent(c, isEnd)
	}
}

func (n *TrieNode) Search(s string) bool {
	return n.checkStringForPresence(s, false)
}

func (n *TrieNode) checkStringForPresence(s string, checkPrefixOnly bool) bool {
	s = strings.ToLower(s)

	t := n
	for _, c := range s {
		charIdx := n.getRelativePosition(c)

		cNode := t.children[charIdx]
		if cNode == nil {
			return false
		}

		t = cNode
	}

	if checkPrefixOnly {
		return true
	} else {
		if t.isEnd {
			return true
		} else {
			return false
		}
	}
}

func (n *TrieNode) StartsWith(s string) bool {
	return n.checkStringForPresence(s, true)
}
