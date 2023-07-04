package Q26_CostOfClimbingStairs

type Prac23MayBtmUp struct {
}

func (s Prac23MayBtmUp) FindMinCostOfClimbingStairs(cost []int) int {
	prev := cost[0]
	latest := cost[1]
	for i := 2; i < len(cost); i++ {
		latest, prev = cost[i]+minVal(latest, prev), latest
	}

	return minVal(prev, latest)
}

//func (s Prac23MayBtmUp) FindMinCostOfClimbingStairs(cost []int) int {
//	n := len(cost)
//	dp := make([]int, n, n)
//
//	dp[0] = cost[0]
//	dp[1] = cost[1]
//	for i := 2; i < n; i++ {
//		dp[i] = cost[i] + minVal(dp[i-1], dp[i-2])
//	}
//
//	//return minVal(minPrev, min), 2
//	return minVal(dp[n-1], dp[n-2]), 2
//}
