package Q26_CostOfClimbingStairs

// start time: 	11:27 pm
// end time: 	11:32 pm

type Prac23June02BtmUpEff struct {
}

func (p Prac23June02BtmUpEff) FindMinCostOfClimbingStairs(cost []int) int {
	old, latest := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		latest, old = cost[i]+p.findMin(latest, old), latest
	}

	return p.findMin(old, latest)
}

func (p Prac23June02BtmUpEff) findMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
