package minTimeToEachNode

import "math"

type Soln1 struct {
}

// using Bellman-Ford Algorithm
func (s Soln1) findMinTime(n int, edges []Node, start int) int {
	const Infinity = math.MaxInt32

	minTimes := make([]int, n, n)
	for i := 0; i < len(minTimes); i++ {
		if i == (start - 1) {
			minTimes[i] = 0 // initial node
			continue
		}

		minTimes[i] = Infinity
	}

	for i := 0; i < n-1; i++ {
		var count int
		for _, e := range edges {
			if minTimes[e.start-1] != Infinity {
				newDist := minTimes[e.start-1] + e.weight
				if newDist < minTimes[e.end-1] {
					minTimes[e.end-1] = newDist
					count++
				}
			}
		}

		if count == 0 {
			break
		}
	}

	var max int
	for _, v := range minTimes {
		if v == Infinity {
			return -1
		}

		if max < v {
			max = v
		}
	}

	return max
}
