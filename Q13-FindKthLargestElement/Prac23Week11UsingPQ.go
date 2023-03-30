package Q13_FindKthLargestElement

import "container/heap"

type Prac23Week11UsingPQ struct {
}

func (p Prac23Week11UsingPQ) findKthLargestElement(arr []int, k int) int {
	h := &PQHeap{}
	heap.Init(h)

	for _, v := range arr {
		heap.Push(h, v)
	}

	var kthVal int
	for i := 0; i < k; i++ {
		kthVal = heap.Pop(h).(int)
	}

	return kthVal
}

type PQHeap []int

func (h PQHeap) Len() int {
	return len(h)
}

func (h PQHeap) Less(i, j int) bool {
	if h[i] > h[j] {
		return true
	} else {
		return false
	}
}

func (h PQHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PQHeap) Push(v any) {
	*h = append(*h, v.(int))
}

func (h *PQHeap) Pop() any {
	origLen := len(*h)
	v := (*h)[origLen-1]
	(*h) = (*h)[:origLen-1]
	return v
}
