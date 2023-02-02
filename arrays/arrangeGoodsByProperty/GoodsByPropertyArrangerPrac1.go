package arrangeGoodsByProperty

import (
	"container/heap"
	_ "container/heap"
)

type Prac1 struct {
	items        []Item
	propertyName string
}

func (h *Prac1) ArrangeGoodsByProperty(items []Item, propName string) []Item {
	h.propertyName = propName
	for _, v := range items {
		heap.Push(h, v)
	}

	arranged := make([]Item, 0, len(items))
	for h.Len() > 0 {
		v := heap.Pop(h)
		arranged = append(arranged, v.(Item))
	}

	return arranged
}

func (h *Prac1) Len() int {
	return len(h.items)
}
func (h *Prac1) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *Prac1) Less(i, j int) bool {
	if h.propertyName == "name" {
		if h.items[i].name < h.items[j].name {
			return true
		} else {
			return false
		}
	} else {
		if h.items[i].expiryTimeInDays < h.items[j].expiryTimeInDays {
			return true
		} else {
			return false
		}
	}
}

func (h *Prac1) Pop() interface{} {
	if h.Len() == 0 {
		return nil
	}

	origLen := len(h.items)
	v := h.items[origLen-1]
	h.items = h.items[:origLen-1]

	return v
}

func (h *Prac1) Push(v interface{}) {
	h.items = append(h.items, v.(Item))
}
