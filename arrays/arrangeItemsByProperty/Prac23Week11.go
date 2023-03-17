package arrangeItemsByProperty

import (
	"container/heap"
	"strings"
)

// start time: 9:13 pm
// end time: 9:31 pm
type Prac23Week11 struct {
}

func (p Prac23Week11) Arrange(items []Item, propName string) []Item {
	h := &PQImplW11{propName: propName}
	for _, v := range items {
		heap.Push(h, v)
	}

	arranged := make([]Item, 0, h.Len())
	for h.Len() > 0 {
		arranged = append(arranged, heap.Pop(h).(Item))
	}

	return arranged
}

type PQImplW11 struct {
	items    []Item
	propName string
}

func (pq PQImplW11) Len() int {
	return len(pq.items)
}

func (pq PQImplW11) Less(i, j int) bool {
	if strings.EqualFold(pq.propName, "name") {
		if pq.items[i].name < pq.items[j].name {
			return true
		} else {
			return false
		}
	} else {
		if pq.items[i].expiryTimeInDays < pq.items[j].expiryTimeInDays {
			return true
		} else {
			return false
		}
	}
}

func (pq PQImplW11) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *PQImplW11) Push(v any) {
	pq.items = append(pq.items, v.(Item))
}

func (pq *PQImplW11) Pop() any {
	origLen := len(pq.items)
	v := pq.items[origLen-1]
	pq.items = pq.items[:origLen-1]
	return v
}
