package Q25_NetworkTimeDelay

import (
	"container/heap"
	"math"
)

type Prac1 struct {
}

type NodeDist struct {
	target int
	dist   int
}

func (p1 Prac1) NetworkDelayTime(n, start int, nodes [][]int) int {
	//prepare the adjacency list
	g := p1.prepareAdjacencyList(n, nodes)
	//apply dijkstra algorithm to find the shortest path
	//create an array of nodes indicating the min distance from the starting position to that node
	dists := make([]int, n, n)
	for i := range dists {
		//initialize it with infinity to indicate that they are not connected in any way
		dists[i] = math.MaxInt32
	}

	finalNodes := make(map[int]bool)

	h := &PQNodeDist{}
	heap.Push(h, NodeDist{start, 0})

	for {
		if h.Len() == 0 {
			break
		}
		src := heap.Pop(h).(NodeDist)
		dists[src.target] = src.dist
		finalNodes[src.target] = true

		//iterate over the immediate neighbors, find the min distances and store it in a priority Q such
		//that it orders PQ's elements based on the distances only.
		adjList := g[src.target]
		for _, v := range adjList {
			if finalNodes[v.target] {
				continue
			}
			v.dist = src.dist + v.dist
			p1.updatePQ(h, v, dists)
		}
	}

	//iterate over the minDist array, check if any of its value is still infinity
	var max int
	for _, v := range dists {
		if v == math.MaxInt32 {
			//which indicates that node is not reachable from the starting point, so return -1
			return -1
		}

		if max < v {
			max = v
		}
	}

	//otherwise find the max node and return its distance
	return max
}

// store tuples: target node index and distance from starting node say TargetDist into this PQ
// init PQ with the starting node and 0 distance
// get the min value from this PQ
// iterate over its adjacent nodes
// for each adjacent node add the new node by adding its parent distance Plus this new distance into PQ
// if it is already present in the PQ, then update its PQ entry if its new distance is less than the one stored in PQ
// otherwise add the new entry
// take care of updating the dist array also
// End of PQ
func (p1 Prac1) updatePQ(h *PQNodeDist, adjNode NodeDist, dists []int) {
	var values []NodeDist

	for h.Len() > 0 {
		node := heap.Pop(h).(NodeDist)
		if node.target == adjNode.target {
			if adjNode.dist < node.dist {
				dists[adjNode.target] = node.dist
			}
			break
		}

		values = append(values, node)
	}

	values = append(values, adjNode)

	for _, v := range values {
		heap.Push(h, v)
	}
}

func (p1 Prac1) prepareAdjacencyList(n int, nodes [][]int) [][]NodeDist {
	g := make([][]NodeDist, n, n)
	for _, v := range nodes {
		adjNodes := g[v[0]]
		if len(g[v[0]]) == 0 {
			adjNodes = make([]NodeDist, 0)
		}

		adjNodes = append(adjNodes, NodeDist{target: v[1], dist: v[2]})
		g[v[0]] = adjNodes
	}

	return g
}

type PQNodeDist []NodeDist

func (pq *PQNodeDist) Len() int {
	return len(*pq)
}

func (pq *PQNodeDist) Swap(i, j int) {
	if i == -1 || j == -1 {
		return
	}
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PQNodeDist) Less(i, j int) bool {
	if (*pq)[i].dist < (*pq)[j].dist {
		return true
	} else {
		return false
	}
}

func (pq *PQNodeDist) Push(v any) {
	*pq = append(*pq, v.(NodeDist))
}

func (pq *PQNodeDist) Pop() any {
	origLen := len(*pq)
	v := (*pq)[origLen-1]
	*pq = (*pq)[:origLen-1]

	return v
}
