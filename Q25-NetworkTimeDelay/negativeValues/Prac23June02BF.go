package negativeValues

import "math"

// Start time: 	12:47 pm
// End time:	01:08 pm

type Prac23June02BF struct {
}

func (p Prac23June02BF) NetworkDelayTime(times [][]int, n, k int) int {
	adjList := p.createAdjList(times, n)

	delays := make([]int, n, n)
	for i := range delays {
		delays[i] = math.MaxInt
	}
	delays[k-1] = 0 //starting node delay is zero

	for i := 0; i < n-1; i++ {
		q := make([]int, 0)
		q = append(q, k-1)

		for len(q) > 0 {
			start := q[0]
			q = q[1:]

			if delays[start] == math.MaxInt {
				continue
			}

			for _, adj := range adjList[start] {
				end := adj[0]
				delay := adj[1]

				d1 := delays[start] + delay
				if d1 < delays[end] {
					delays[end] = d1
					q = append(q, end)
				}
			}
		}
	}

	var max int
	for i := range delays {
		if delays[i] == math.MaxInt {
			return -1
		}

		if max < delays[i] {
			max = delays[i]
		}
	}

	return max
}

func (p Prac23June02BF) createAdjList(times [][]int, n int) [][][]int {
	adjList := make([][][]int, n, n)

	for _, t := range times {
		start := t[0] - 1
		end := t[1] - 1
		delay := t[2]

		adjs := adjList[start]
		if len(adjs) == 0 {
			adjs = make([][]int, 0)
		}

		adjList[start] = append(adjs, []int{end, delay})
	}

	return adjList
}
