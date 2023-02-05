package findKthLargestElement

import "container/heap"

type Prac1 struct {
	arr []int
}

func (p *Prac1) findKthLargestElement(arr []int, k int) int {
	if len(arr) == 0 {
		return 0
	}

	p.arr = p.arr[0:0]

	//freqByElement := make(map[int]int, 0)
	for _, v := range arr {
		heap.Push(p, v)
		//if i, ok := freqByElement[v]; !ok {
		//	freqByElement[v] = 1
		//} else {
		//	freqByElement[v] = i + 1
		//}
	}

	//for k := range freqByElement {
	//	heap.Push(p, k)
	//}
	// discard top k-1 elements
	i := 1
	//var val int
	for i < k {
		//val := heap.Pop(p).(int)
		heap.Pop(p)
		//if f, ok := freqByElement[val]; !ok {
		i++
		//} else {
		//	i += f
		//}
	}

	//if i > k {
	//	return val
	//}

	return heap.Pop(p).(int) // kth element
}

func (p *Prac1) Len() int {
	return len(p.arr)
}

func (p *Prac1) Less(i, j int) bool {
	return p.arr[i] >= p.arr[j]
}

func (p *Prac1) Swap(i, j int) {
	p.arr[i], p.arr[j] = p.arr[j], p.arr[i]
}

func (p *Prac1) Push(v interface{}) {
	p.arr = append(p.arr, v.(int))
}

func (p *Prac1) Pop() interface{} {
	origLen := len(p.arr)
	v := p.arr[origLen-1]
	p.arr = p.arr[:origLen-1]
	return v
}
