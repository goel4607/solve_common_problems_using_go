package Q25_NetworkTimeDelay

import "math"
import "container/heap"

type Prac23AprDijk struct {
}

func (s Prac23AprDijk) NetworkDelayTime(times [][]int, n, k int) int {
	//create an adjacency list graph
	g := s.createAdjList(n, times)
	//create an arr of min distances from k node with initial distances as infinite to indicate that it is not
	dists := make([]int, n, n)
	for i := range dists {
		dists[i] = math.MaxInt
	}

	dists[k-1] = 0 //since the distance from k to itself is zero

	//create a priority queue that uses the distances to consult
	h := MyHeap{vertexes: make([]int, 0), dists: &dists}
	heap.Init(&h)

	heap.Push(&h, k-1)

	for h.Len() > 0 {
		currIdx := heap.Pop(&h).(int)
		adjList := g[currIdx]
		for _, adj := range adjList {
			nbrVertex := adj[0]
			nbrDist := adj[1]

			newDist := dists[currIdx] + nbrDist
			if newDist < dists[nbrVertex] {
				dists[nbrVertex] = newDist
				heap.Push(&h, nbrVertex)
			}
		}
	}

	var max int
	for i := range dists {
		if max < dists[i] {
			max = dists[i]
		}
	}

	if max == math.MaxInt { //disconnected graph
		return -1
	}

	return max
}

func (s Prac23AprDijk) createAdjList(n int, times [][]int) [][][]int {
	g := make([][][]int, n, n)
	for _, t := range times {
		start := t[0] - 1
		end := t[1] - 1

		adj := make([]int, 0, 0)
		adj = append(adj, end, t[2])

		if g[start] == nil {
			adjList := make([][]int, 0)
			adjList = append(adjList, adj)

			g[start] = adjList
		} else {
			adjList := g[start]
			adjList = append(adjList, adj)
			g[start] = adjList
		}
	}

	return g
}

type MyHeap struct {
	vertexes []int
	dists    *[]int
}

func (h *MyHeap) Len() int {
	return len(h.vertexes)
}

func (h *MyHeap) Swap(i, j int) {
	h.vertexes[i], h.vertexes[j] = h.vertexes[j], h.vertexes[i]
}

func (h *MyHeap) Less(i, j int) bool {
	if (*h.dists)[h.vertexes[i]] < (*h.dists)[h.vertexes[j]] {
		return true
	}

	return false
}

func (h *MyHeap) Push(v any) {
	val := v.(int)
	h.vertexes = append(h.vertexes, val)
}

func (h *MyHeap) Pop() any {
	oldLen := h.Len()
	v := h.vertexes[oldLen-1]
	h.vertexes = h.vertexes[:oldLen-1]
	return v
}
