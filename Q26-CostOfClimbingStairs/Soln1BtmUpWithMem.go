package Q26_CostOfClimbingStairs

import "math"

type Soln1BtmUpWithMem struct {
}

func (s Soln1BtmUpWithMem) FindMinCostOfClimbingStairs(cost []int) int {
	n := len(cost)

	secLast := cost[0]
	last := cost[1]

	var numCalls int
	for i := 2; i < n; i++ {
		numCalls++
		tmp := cost[i] + int(math.Min(float64(last), float64(secLast)))
		secLast = last
		last = tmp
	}

	var min int
	if last > secLast {
		min = secLast
	} else {
		min = last
	}

	return min
}
