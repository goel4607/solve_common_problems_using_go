package minCostOfClimbingStairs

import "math"

type Soln1 struct {
	numCalls int
}

func (s *Soln1) findMinCostOfClimbingStairs(cost []int) (int, int) {
	n := len(cost)
	mem := make(map[int]int)

	return int(math.Min(float64(s.findMinCost(cost, n-1, mem)), float64(s.findMinCost(cost, n-2, mem)))), s.numCalls
}

func (s *Soln1) findMinCost(cost []int, n int, mem map[int]int) int {
	s.numCalls++

	if v, ok := mem[n]; ok {
		return v
	}

	if n <= 1 {
		mem[n] = cost[n]
		return cost[n]
	}

	min := cost[n] + int(math.Min(float64(s.findMinCost(cost, n-1, mem)), float64(s.findMinCost(cost, n-2, mem))))
	mem[n] = min
	return min
}
