package Q26_CostOfClimbingStairs

import "math"

type Soln1MemUsingArr struct {
	numCalls int
}

func (s *Soln1MemUsingArr) findMinCostOfClimbingStairs(cost []int) (int, int) {
	n := len(cost)
	mem := make([]int, n, n)
	for i := 0; i < len(mem); i++ {
		mem[i] = -1
	}

	return int(math.Min(float64(s.findMinCost(cost, n-1, mem)), float64(s.findMinCost(cost, n-2, mem)))), s.numCalls
}

func (s *Soln1MemUsingArr) findMinCost(cost []int, n int, mem []int) int {
	s.numCalls++

	if mem[n] != -1 {
		return mem[n]
	}

	if n <= 1 {
		mem[n] = cost[n]
		return cost[n]
	}

	mem[n] = cost[n] + int(math.Min(float64(s.findMinCost(cost, n-1, mem)), float64(s.findMinCost(cost, n-2, mem))))
	return mem[n]
}
