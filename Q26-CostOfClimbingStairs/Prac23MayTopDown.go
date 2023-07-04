package Q26_CostOfClimbingStairs

type Prac23MayTry1 struct {
}

// First Step: Establish a reoccurrence relation: minCost[i] = cost[i] + min(minCost[i-1], minCost[i-2])
// Second Step: Create the initial recursive function and derive the space and time efficiency
// Third Step: Efficient solution using either memoization or something else!! so that we can cut-down on un-necessary cycles
func (s Prac23MayTry1) FindMinCostOfClimbingStairs(cost []int) int {
	minCostVals := make([]int, len(cost), len(cost))
	for i := range cost {
		minCostVals[i] = -1
	}
	return s.minCost(cost, minCostVals, len(cost), 1)
}

func (s Prac23MayTry1) minCost(cost []int, minCostVals []int, i, numCalls int) int {
	if i <= 1 {
		return cost[i]
	}

	if i < len(cost) && minCostVals[i] != -1 {
		return minCostVals[i]
	}

	var prev, nc1 int
	if minCostVals[i-1] == -1 {
		prev = s.minCost(cost, minCostVals, i-1, numCalls+1)
		minCostVals[i-1] = prev
	} else {
		prev = minCostVals[i-1]
	}

	var nxtToPrev, nc2 int
	if minCostVals[i-2] == -1 {
		nxtToPrev = s.minCost(cost, minCostVals, i-2, numCalls+1)
		minCostVals[i-2] = prev
	} else {
		nxtToPrev = minCostVals[i-2]
	}

	min := minVal(prev, nxtToPrev)
	numCalls = minVal(nc1, nc2)

	if i < len(cost) {
		return min + cost[i]
	} else {
		return min
	}
}

func minVal(a int, b int) int {
	if a < b {
		return a
	}

	return b
}
