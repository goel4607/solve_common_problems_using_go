package Q25_NetworkTimeDelay

import "math"

// start time: 10:30 pm
// end time: 11:20 pm
type Prac23June01Dijk struct {
}

func (p Prac23June01Dijk) NetworkDelayTime(times [][]int, n, k int) int {

	weights := make([]int, n, n)
	for i := 0; i < len(weights); i++ {
		weights[i] = math.MaxInt
	}

	//distance of starting node to itself is zero
	weights[k-1] = 0

	adjList := p.createAdjList(n, times)

	q := make([]int, 0)
	q = append(q, k-1) //add starting node to Q

	excludedNodes := make(map[int]bool, 0)

	for len(q) > 0 {
		startNode := q[0]
		q = q[1:] //remove first element

		for _, adj := range adjList[startNode] {
			num := adj[0]
			wt := adj[1]

			newWt := weights[startNode] + wt
			if newWt < weights[num] {
				weights[num] = newWt
			}
		}

		excludedNodes[startNode] = true

		minWtNode := p.findMinWeightNodeId(weights, excludedNodes)
		if minWtNode == -1 {
			break
		}
		q = append(q, minWtNode)
	}

	maxWt := -1
	for _, wt := range weights {
		if wt == math.MaxInt {
			return -1
		}

		if maxWt < wt {
			maxWt = wt
		}
	}

	return maxWt
}

func (p Prac23June01Dijk) findMinWeightNodeId(wts []int, excludedNodes map[int]bool) int {
	minWt := math.MaxInt
	minWtId := -1
	for i, wt := range wts {
		if excludedNodes[i] {
			continue
		}

		if minWt > wt {
			minWt = wt
			minWtId = i
		}
	}

	return minWtId
}

func (p Prac23June01Dijk) createAdjList(n int, times [][]int) [][][]int {
	adjList := make([][][]int, n, n)

	for _, t := range times {
		start := t[0] - 1
		nws := adjList[start]
		if len(nws) == 0 {
			nws = make([][]int, 0)
		}

		adjList[start] = append(nws, []int{t[1] - 1, t[2]})
	}

	return adjList
}
