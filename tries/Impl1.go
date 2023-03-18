package tries

import "strings"

const (
	NumAlphabets  = 26
	StartingIndex = rune('a')
)

type TrieNode struct {
	C        rune
	IsEnd    bool
	Children [NumAlphabets]*TrieNode
}

func (n *TrieNode) getRelativePosition(c rune) int {
	return int(c - StartingIndex)
}

func (n *TrieNode) addChar(c rune, isEnding bool) *TrieNode {
	arrIdx := n.getRelativePosition(c)
	newChar := &TrieNode{
		C:        c,
		IsEnd:    isEnding,
		Children: [26]*TrieNode{},
	}

	n.Children[arrIdx] = newChar

	return newChar
}

func (n *TrieNode) addCharIfNotPresent(c rune, isEnding bool) *TrieNode {
	arrIdx := n.getRelativePosition(c)

	if n.Children[arrIdx] == nil {
		newChar := &TrieNode{
			C:     c,
			IsEnd: isEnding,
		}

		n.Children[arrIdx] = newChar
		return newChar
	}

	return n.Children[arrIdx]
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

		cNode := t.Children[charIdx]
		if cNode == nil {
			return false
		}

		t = cNode
	}

	if checkPrefixOnly {
		return true
	} else {
		if t.IsEnd {
			return true
		} else {
			return false
		}
	}
}

func (n *TrieNode) StartsWith(s string) bool {
	return n.checkStringForPresence(s, true)
}
