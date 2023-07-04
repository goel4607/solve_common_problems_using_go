package Q23_TotalMinFinderToInformAllEmps

// start time: 	4:58 pm
// end time:

type Prac23June01BFS struct {
}

func (p Prac23June01BFS) numOfMinutes(n int, headId int, mgrs, informTime []int) int {
	adjList := p.createAdjList(mgrs)

	q := make([]int, 0)
	q = append(q, mgrs[headId])

	var totalTime int
	for len(q) > 0 {
		m := q[0]
		q = q[1:]

		directReports := adjList[m]
		if len(directReports) == 0 {
			continue
		}

		_ = len(directReports)

	}

	return totalTime
}

func (p Prac23June01BFS) createAdjList(mgrs []int) [][]int {
	adjList := make([][]int, 0)

	for e := 0; e < len(mgrs); e++ {
		m := mgrs[e]
		if m == -1 { // head does not have any manager
			continue
		}

		directReports := adjList[m]
		if len(directReports) == 0 {
			directReports = make([]int, 0)
		}

		adjList[m] = append(directReports, e)
	}

	return adjList
}
