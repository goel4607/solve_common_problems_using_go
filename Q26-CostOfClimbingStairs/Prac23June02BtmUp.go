package Q26_CostOfClimbingStairs

// start time: 	10:58 pm
// end time: 	11:27 pm

type Prac23June02BtmUp struct {
}

func (p Prac23June02BtmUp) FindMinCostOfClimbingStairs(cost []int) int {
	n := len(cost)
	cache := make([]int, n, n)
	for i := 0; i < n; i++ {
		if i <= 1 {
			cache[i] = cost[i]
		} else {
			cache[i] = cost[i] + p.getMin(cache[i-1], cache[i-2])
		}
	}

	return p.getMin(cache[n-1], cache[n-2])
}

func (p Prac23June02BtmUp) getMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}
