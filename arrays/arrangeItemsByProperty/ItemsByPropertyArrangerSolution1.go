package arrangeItemsByProperty

import (
	"container/heap"
	"fmt"
)

type ItemsByPropertyArrangerSolution1 struct {
	items    []Item
	property string
}

func (h *ItemsByPropertyArrangerSolution1) Arrange(items []Item, propName string) []Item {
	h.property = propName
	for _, v := range items {
		heap.Push(h, v)
	}

	arranged := make([]Item, 0, len(items))
	for h.Len() > 0 {
		v := heap.Pop(h)
		if val, ok := v.(Item); ok {
			arranged = append(arranged, val)
		}
	}

	return arranged
}

func (h *ItemsByPropertyArrangerSolution1) Len() int {
	return len(h.items)
}

func (h *ItemsByPropertyArrangerSolution1) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *ItemsByPropertyArrangerSolution1) Less(i, j int) bool {
	if h.property == "name" {
		return h.items[i].name < h.items[j].name
	} else {
		return h.items[i].expiryTimeInDays < h.items[j].expiryTimeInDays
	}
}

func (h *ItemsByPropertyArrangerSolution1) Push(v interface{}) {
	h.items = append(h.items, v.(Item))
}

func (h *ItemsByPropertyArrangerSolution1) Pop() interface{} {
	origItems := h.items
	origLen := len(h.items)

	v := origItems[origLen-1]
	h.items = origItems[:origLen-1]

	return v
}

func (h *ItemsByPropertyArrangerSolution1) String() string {
	s := ""
	for _, v := range h.items {
		s += fmt.Sprintf("[%v, %v]", v.name, v.expiryTimeInDays)
	}

	return s
}
