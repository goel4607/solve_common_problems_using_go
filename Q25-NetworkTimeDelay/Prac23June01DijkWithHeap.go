package Q25_NetworkTimeDelay

import "math"
import "container/heap"

// start time: 12:41 am
// end   time: 1:38 am

type Prac23June01DijkWithHeap struct {
}

func (p Prac23June01DijkWithHeap) NetworkDelayTime(times [][]int, n, k int) int {
	dists := make([]int, n)
	for i := 0; i < n; i++ {
		dists[i] = math.MaxInt
	}
	dists[k-1] = 0

	h := &Prac23June01Heap{
		vertices: make([]int, 0),
		dists:    &dists,
	}
	heap.Init(h)

	adjList := p.createAdjList(times, n)

	heap.Push(h, k-1)

	for h.Len() > 0 {
		minDist := heap.Pop(h).(int)

		for _, adjs := range adjList[minDist] {
			vtx := adjs[0]
			wt := adjs[1]

			newWt := dists[minDist] + wt
			if newWt < dists[vtx] {
				dists[vtx] = newWt

				heap.Push(h, vtx)
			}
		}
	}

	var max int
	for _, d := range dists {
		if max < d {
			max = d
		}
	}

	if max == math.MaxInt { //disconnected graph or start vertex cannot reach all the vertexes
		return -1
	}

	return max
}

func (p Prac23June01DijkWithHeap) createAdjList(times [][]int, n int) [][][]int {
	adjList := make([][][]int, n, n)

	for _, t := range times {
		start := t[0] - 1
		end := t[1] - 1
		w := t[2]

		adjs := adjList[start]
		if len(adjs) == 0 {
			adjs = make([][]int, 0)
		}

		adjList[start] = append(adjs, []int{end, w})
	}

	return adjList
}

type Prac23June01Heap struct {
	vertices []int
	dists    *[]int
}

func (h *Prac23June01Heap) Len() int {
	return len(h.vertices)
}

func (h *Prac23June01Heap) Swap(i, j int) {
	h.vertices[i], h.vertices[j] = h.vertices[j], h.vertices[i]
}

func (h *Prac23June01Heap) Less(i, j int) bool {
	if (*h.dists)[h.vertices[i]] < (*h.dists)[h.vertices[j]] {
		return true
	}

	return false
}

func (h *Prac23June01Heap) Push(i any) {
	h.vertices = append(h.vertices, i.(int))
}

func (h *Prac23June01Heap) Pop() any {
	oldLen := len(h.vertices)
	v := h.vertices[oldLen-1]
	h.vertices = h.vertices[:oldLen-1]
	return v
}
