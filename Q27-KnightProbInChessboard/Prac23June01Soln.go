package minCostOfClimbingStairs

// start time: 	11:42 pm
// end time:	12:16 pm

type Prac23June01Soln struct {
}

type Move struct {
	r int
	c int
}

var AllMoves = [8]Move{
	0: {-2, 1},
	1: {-1, 2},
	2: {1, 2},
	3: {2, 1},
	4: {2, -1},
	5: {1, -2},
	6: {-1, -2},
	7: {-2, -1},
}

func (p Prac23June01Soln) computeProbabilityThatKnightIsStillOnCB(n int, k int, r, c int) float64 {
	cache := make([][][]float64, k, k)
	for i := range cache {
		cache[i] = make([][]float64, n, n)
		for j := range cache[i] {
			cache[i][j] = make([]float64, n, n)
			for m := range cache[i][j] {
				cache[i][j][m] = -1
			}
		}
	}

	return p.computeProbAtStepK(n, k, r, c, cache)
}

func (p Prac23June01Soln) computeProbAtStepK(n, k, r, c int, cache [][][]float64) float64 {
	if r < 0 || r >= n || c < 0 || c >= n {
		return 0
	}

	if k == 0 {
		return float64(1)
	}

	if cache[k-1][r][c] != -1 {
		return cache[k-1][r][c]
	}

	var prob float64
	for _, m := range AllMoves {
		prob += p.computeProbAtStepK(n, k-1, r+m.r, c+m.c, cache)
	}

	prob = prob / float64(len(AllMoves))

	cache[k-1][r][c] = prob

	return prob
}
