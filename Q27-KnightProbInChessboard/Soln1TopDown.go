package minCostOfClimbingStairs

type Soln1TopDown struct {
}

type Position struct {
	r int
	c int
}

var Moves = [8]Position{
	0: {-2, 1},
	1: {-1, 2},
	2: {1, 2},
	3: {2, 1},
	4: {2, -1},
	5: {1, -2},
	6: {-1, -2},
	7: {-2, -1},
}

func (s Soln1TopDown) computeProbabilityThatKnightIsStillOnCB(n, k, r, c int) float64 {
	dp := make([][][]float64, k, k)
	for i := range dp {
		dp[i] = make([][]float64, n, n)
		for j := range dp[i] {
			dp[i][j] = make([]float64, n, n)
			for m := range dp[i][j] {
				dp[i][j][m] = -1
			}
		}
	}

	return s.computeProbAtPosition(n, k, r, c, dp)
}

func (s Soln1TopDown) computeProbAtPosition(n, k, r, c int, dp [][][]float64) float64 {
	if r < 0 || r >= n || c < 0 || c >= n {
		return 0
	}

	if k == 0 {
		return 1
	}

	if dp[k-1][r][c] != -1 {
		return dp[k-1][r][c]
	}

	var prob float64
	for _, m := range Moves {
		p := s.computeProbAtPosition(n, k-1, r+m.r, c+m.c, dp)
		prob += p
	}

	prob = prob / float64(len(Moves))

	dp[k-1][r][c] = prob

	return prob
}
