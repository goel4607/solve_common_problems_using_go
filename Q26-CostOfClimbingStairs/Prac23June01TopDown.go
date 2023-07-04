package Q26_CostOfClimbingStairs

type Prac23June01TopDown struct {
}

func (p Prac23June01TopDown) FindMinCostOfClimbingStairs(cost []int) int {
	total := len(cost)
	if total == 0 {
		return -1
	}

	cache := make(map[int]int)

	return findMin(p.findMinCost(cost, total-1, cache), p.findMinCost(cost, total-2, cache))
}

// formulae: minCost = cost(n) + min(cost(n-1), cost(cost-2)

func (p Prac23June01TopDown) findMinCost(cost []int, step int, cache map[int]int) int {
	if step <= 1 {
		return cost[step]
	}

	if v, ok := cache[step]; ok {
		return v
	}

	c1 := p.findMinCost(cost, step-1, cache)
	cache[step-1] = c1

	c2 := p.findMinCost(cost, step-2, cache)
	cache[step-2] = c2

	return cost[step] + findMin(c1, c2)
}

func findMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}
