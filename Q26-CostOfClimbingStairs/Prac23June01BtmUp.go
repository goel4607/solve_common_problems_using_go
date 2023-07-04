package Q26_CostOfClimbingStairs

type Prac23June01BtmUp struct {
}

func (p Prac23June01BtmUp) FindMinCostOfClimbingStairs(cost []int) int {
	mem := make(map[int]int, 0)
	n := len(cost)
	for i := 0; i < n; i++ {
		if i <= 1 {
			mem[i] = cost[i]
		} else {
			mem[i] = cost[i] + p.findMin(mem[i-1], mem[i-2])
		}
	}

	return p.findMin(mem[n-1], mem[n-2])
}

func (p Prac23June01BtmUp) findMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}
