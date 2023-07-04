package negativeValues

import "math"

type Prac23June01BF struct {
}

func (p Prac23June01BF) NetworkDelayTime(times [][]int, n, k int) int {
	dists := make([]int, n, n)
	for i := 0; i < len(dists); i++ {
		dists[i] = math.MaxInt
	}

	dists[k-1] = 0

	for i := 0; i < n; i++ {
		var hasChanged bool

		for _, t := range times {
			start := t[0] - 1
			end := t[1] - 1
			wt := t[2]

			if dists[start] == math.MaxInt {
				continue
			}

			newWt := dists[start] + wt
			if newWt < dists[end] {
				dists[end] = newWt
				hasChanged = true
			}
		}

		if !hasChanged {
			break
		}
	}

	var max int
	for _, v := range dists {
		if max < v {
			max = v
		}
	}

	if max == math.MaxInt {
		return -1
	}

	return max
}
