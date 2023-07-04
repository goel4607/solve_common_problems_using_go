package Q26_CostOfClimbingStairs

// start time: 	10:19 pm
// end time:	10:45 pm

type Prac23June02TopDown struct {
}

func (p Prac23June02TopDown) FindMinCostOfClimbingStairs(cost []int) int {
	//create a formula: minCost = cost + min(minCostAtStep[n-1], minCostAtStep[n-2])
	//declare a cache array
	n := len(cost)
	cache := make([]int, n, n)
	//start from the top most, since its cost is zero
	return 0 + p.getMin(p.findMinCost(cost, cache, n-1), p.findMinCost(cost, cache, n-2))
}

func (p Prac23June02TopDown) findMinCost(cost []int, dp []int, step int) int {
	if step <= 1 {
		dp[step] = cost[step]
		return dp[step]
	}

	if dp[step] != 0 {
		return dp[step]
	}

	c1 := p.findMinCost(cost, dp, step-1)
	dp[step-1] = c1

	c2 := p.findMinCost(cost, dp, step-2)
	dp[step-2] = c2

	return p.getMin(c1, c2) + cost[step]
}

func (p Prac23June02TopDown) getMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}
