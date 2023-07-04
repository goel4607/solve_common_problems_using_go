package Q26_CostOfClimbingStairs

import "math"

type Soln1WithoutMemoization struct {
	numCalls int
}

func (s *Soln1WithoutMemoization) FindMinCostOfClimbingStairs(cost []int) int {
	n := len(cost)

	return int(math.Min(float64(s.findMinCost(cost, n-1)), float64(s.findMinCost(cost, n-2))))
	//return s.findMinCost(cost, len(cost), make(map[int]int)), s.numCalls
}

func (s *Soln1WithoutMemoization) findMinCost(cost []int, n int) int {
	s.numCalls++

	if n <= 1 {
		return cost[n]
	}

	return cost[n] + int(math.Min(float64(s.findMinCost(cost, n-1)), float64(s.findMinCost(cost, n-2))))
}
