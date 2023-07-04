package negativeValues

import "math"

// start time: 	3:16 pm
// end time:

type Prac23July01BellmanFord struct {
}

func (p Prac23July01BellmanFord) NetworkDelayTime(times [][]int, n, k int) int {
	adjList := p.createAdjacencyList(times, n)

	distances := make([]int, n, n)
	for i := range distances {
		distances[i] = math.MaxInt
	}
	distances[k-1] = 0

	for i := 0; i < n-1; i++ {
		q := make([]int, 0)
		q = append(q, k-1)

		for len(q) > 0 {
			src := q[0]
			q = q[1:]

			if distances[src] == math.MaxInt {
				continue
			}

			//var isNewEntry bool
			for _, adj := range adjList[src] {
				target := adj[0]
				dist := adj[1]

				newDist := distances[src] + dist
				if newDist < distances[target] {
					distances[target] = newDist

					q = append(q, target)
					//isNewEntry = true
				}
			}

			//if !isNewEntry { // no new entries found so distances are already optimal
			//	break
			//}
		}
	}

	var max int
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

func (p Prac23July01BellmanFord) createAdjacencyList(times [][]int, n int) [][][]int {
	g := make([][][]int, n, n)
	for _, t := range times {
		src := t[0] - 1
		tgt := t[1] - 1
		dist := t[2]

		adjacentNodes := g[src]
		if len(adjacentNodes) == 0 {
			adjacentNodes = make([][]int, 0)
		}

		g[src] = append(adjacentNodes, []int{tgt, dist})
	}

	return g
}
