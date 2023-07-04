package Q26_CostOfClimbingStairs

type Prac23June01BtmUpEff struct {
}

func (p Prac23June01BtmUpEff) FindMinCostOfClimbingStairs(cost []int) int {
	a, b := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		b, a = cost[i]+p.findMin(a, b), b
	}

	return p.findMin(a, b)
}

func (p Prac23June01BtmUpEff) findMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}
