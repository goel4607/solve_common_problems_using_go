package minTimeToEachNode

import (
	"container/heap"
	"math"
)

type Soln1 struct {
}

func (s Soln1) findMinTime(n, start int, nodes []Node) int {
	g := s.createAdjList(n, nodes)

	minWeights := make([]int, n, n)
	//set the max value for each weight
	for i := range minWeights {
		if i == start {
			minWeights[i] = 0 //weight to reach itself
			continue
		}
		minWeights[i] = math.MaxInt32
	}

	finalNodes := make([]bool, n, n)

	h := &NodeWeightsHeap{}
	heap.Push(h, NodeWeight{index: start, weight: 0})

	for {
		if h.Len() == 0 {
			break
		}

		prev := heap.Pop(h).(NodeWeight)
		finalNodes[prev.index] = true
		minWeights[prev.index] = prev.weight

		list := g[prev.index]
		for _, v := range list {
			if finalNodes[v.index] {
				continue //ignore the finalized node indexes
			}

			v.weight = v.weight + prev.weight // add the weight till the current node
			updateValueIfPresent(h, v)
		}
	}

	max := 0 //we want the longest time
	for _, w := range minWeights {
		if w == math.MaxInt32 { //not visited entry
			return -1
		}

		if max < w {
			max = w
		}
	}

	return max
}

func updateValueIfPresent(h *NodeWeightsHeap, nw NodeWeight) {
	vals := make([]NodeWeight, 0)
	var found bool
	for h.Len() > 0 {
		v := heap.Pop(h).(NodeWeight)
		if v.index == nw.index {
			if v.weight > nw.weight {
				vals = append(vals, nw)
			}
			found = true
			break
		}

		vals = append(vals, v)
	}

	if !found {
		vals = append(vals, nw)
	}

	for _, v := range vals {
		heap.Push(h, v)
	}
}

type NodeWeight struct {
	index  int
	weight int
}

type NodeWeightsHeap []NodeWeight

func (h NodeWeightsHeap) Len() int {
	return len(h)
}

func (h NodeWeightsHeap) Less(i, j int) bool {
	if h[i].weight < h[j].weight {
		return true
	} else {
		return false
	}
}

func (h NodeWeightsHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeWeightsHeap) Push(v any) {
	*h = append(*h, v.(NodeWeight))
}

func (h *NodeWeightsHeap) Pop() any {
	origLen := h.Len()
	v := (*h)[origLen-1]
	*h = (*h)[:origLen-1]

	return v
}

// here each node has 3 values:
// [0] -> start node,
// [1] -> end node,
// [2] -> weight
func (s Soln1) createAdjList(n int, nodes []Node) [][]NodeWeight {
	g := make([][]NodeWeight, n, n)

	for _, v := range nodes {

		var adjList []NodeWeight
		if len(g[v.start]) == 0 {
			adjList = make([]NodeWeight, 0)
		} else {
			adjList = g[v.start]
		}

		adjList = append(adjList, NodeWeight{index: v.end, weight: v.weight})

		g[v.start] = adjList
	}

	return g
}
