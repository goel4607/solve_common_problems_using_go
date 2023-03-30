package Q26_CostOfClimbingStairs

import "math"

type Soln1BtmUp struct {
}

func (s Soln1BtmUp) findMinCostOfClimbingStairs(cost []int) (int, int) {
	var numCalls int

	n := len(cost)

	dp := make([]int, n, n)
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < len(cost); i++ {
		numCalls++

		dp[i] = cost[i] + int(math.Min(float64(dp[i-1]), float64(dp[i-2])))
	}

	var min int
	if dp[n-1] < dp[n-2] {
		min = dp[n-1]
	} else {
		min = dp[n-2]
	}
	return min, numCalls
}
