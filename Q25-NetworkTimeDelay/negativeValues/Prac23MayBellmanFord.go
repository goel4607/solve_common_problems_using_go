package negativeValues

import (
	"math"
)

type Prac23MayBellmanFord struct {
}

func (s Prac23MayBellmanFord) NetworkDelayTime(times [][]int, n, k int) int {
	//init the dist from start node to all nodes with infinity
	dists := make([]int, n, n)
	for i := range dists {
		dists[i] = math.MaxInt
	}

	//set the distance from start to itself as zero
	dists[k] = 0

	//iterate over all the edges n-1 times such that each time we update the distance at that edge
	//this uses dynamic programming memoization technique where we compute each node's distance to the other node
	//utilizing its already computed distance (using memoization)
	//we will iterate only n-1 times since there are n nodes and we are starting from one of the nodes so we need to
	//compute only n - 1 times
	//also there can be negative cycles i.e. certain values where the distances keep reducing even after n-1 times
	for i := 0; i < n-1; i++ {
		var hasChanged bool
		for _, e := range times {
			if dists[e[0]] == math.MaxInt {
				continue //ignore this value since it is already infinity
			}

			d := dists[e[0]] + e[2]
			if d < dists[e[1]] {
				dists[e[1]] = d //update the distance if its new distance is < than the old distance
				hasChanged = true
			}
		}

		if !hasChanged {
			break //no values have changed so there is no use iterating again since nodes are already at min values
		}
	}

	var max int
	for _, d := range dists {
		if max < d {
			max = d
		}
	}

	if max == math.MaxInt {
		return -1
	}

	return max
}
