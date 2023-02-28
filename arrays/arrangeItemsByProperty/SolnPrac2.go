package arrangeItemsByProperty

import "container/heap"

type SolnPrac2 struct {
	items    []Item
	property string
}

func (h *SolnPrac2) Arrange(items []Item, propName string) []Item {
	h.items = items
	h.property = propName

	heap.Init(h)

	orderedItems := make([]Item, 0, len(items))
	for h.Len() > 0 {
		v := heap.Pop(h)
		if i, ok := v.(Item); ok {
			orderedItems = append(orderedItems, i)
		}
	}

	return orderedItems
}

func (h *SolnPrac2) Len() int {
	return len(h.items)
}

func (h *SolnPrac2) Less(i, j int) bool {
	if h.property == "expiryTimeInDays" || len(h.property) == 0 {
		if h.items[i].expiryTimeInDays < h.items[j].expiryTimeInDays {
			return true
		} else {
			return false
		}
	} else if h.property == "name" {
		if h.items[i].name < h.items[j].name {
			return true
		} else {
			return false
		}
	}

	return false
}

func (h *SolnPrac2) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *SolnPrac2) Push(x any) {
	h.items = append(h.items, x.(Item))
}

func (h *SolnPrac2) Pop() any {
	origLen := len(h.items)
	v := h.items[origLen-1]
	h.items = h.items[0 : origLen-1]
	return v
}
