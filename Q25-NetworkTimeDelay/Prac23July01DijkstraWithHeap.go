package Q25_NetworkTimeDelay

import "math"
import "container/heap"

// start time: 	1:40 pm
// end time:	2:40 pm

type Prac23July01DijkstraWithHeap struct {
}

func (p Prac23July01DijkstraWithHeap) NetworkDelayTime(times [][]int, n, k int) int {
	//problem: the problem is to find the max time it takes from the starting node say k to reach out to any node.
	//if any of the node is un-reachable from the starting position, then we return -1

	//solution: create an n node array and init it with Max number. This array represents the distance from the starting
	//node to all the other nodes
	distances := make([]int, n, n)
	for i := 0; i < len(distances); i++ {
		distances[i] = math.MaxInt
	}
	distances[k-1] = 0

	//Then we create an adjacency list representing the nodes adjacent to each node like it would contain an 2D arrary
	//for each node say for k node it would contain all the reachable nodes with each node representing its index at
	//position 0 and distance at position 1
	adjacencyList := p.createAdjacencyList(times, n)

	pq := &Prac23July01Heap{vertices: make([]int, 0), distances: &distances}
	heap.Init(pq)

	//init a Priority Q with K as the beginning position and iterate over the Q till there is no element present in it.
	heap.Push(pq, k-1)

	//get the min element from the PQ and iterate over its adjacent nodes and find the adjacent node distance from this
	//node. Check if its distance is less than the one present in the distances, if yes, then update the dist array with
	//the new distance. Similarly, keep iterating over PQ till there is no element left
	for pq.Len() > 0 {
		v := heap.Pop(pq).(int)

		for _, adjNode := range adjacencyList[v] {
			tgt := adjNode[0]
			dist := adjNode[1]

			newDist := distances[v] + dist
			if newDist < distances[tgt] {
				distances[tgt] = newDist
				heap.Push(pq, tgt)
			}
		}
	}

	//now get the max out of distances and return that value. Also check if any of the node is still un-reachable then
	max := 0
	for _, d := range distances {
		if d == math.MaxInt {
			return -1
		}

		if max < d {
			max = d
		}
	}

	return max
}

func (p Prac23July01DijkstraWithHeap) createAdjacencyList(times [][]int, n int) [][][]int {
	g := make([][][]int, n, n)
	for _, t := range times {
		src := t[0] - 1
		dst := t[1] - 1
		dist := t[2]

		adjList := g[src]
		if len(adjList) == 0 {
			adjList = make([][]int, 0)
		}

		g[src] = append(adjList, []int{dst, dist})
	}

	return g
}

type Prac23July01Heap struct {
	vertices  []int
	distances *[]int
}

func (h *Prac23July01Heap) Len() int {
	return len(h.vertices)
}

func (h *Prac23July01Heap) Swap(i, j int) {
	h.vertices[i], h.vertices[j] = h.vertices[j], h.vertices[i]
}

func (h *Prac23July01Heap) Less(i, j int) bool {
	return (*h.distances)[h.vertices[i]] < (*h.distances)[h.vertices[j]]
}

func (h *Prac23July01Heap) Push(v any) {
	h.vertices = append(h.vertices, v.(int))
}

func (h *Prac23July01Heap) Pop() any {
	oldLen := len(h.vertices)
	v := h.vertices[oldLen-1] // last element
	h.vertices = h.vertices[0 : oldLen-1]

	return v
}
