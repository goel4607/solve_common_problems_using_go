package Q23_TotalMinFinderToInformAllEmps

// start time: 4:01 pm
// end time:   4:22 pm
type Prac23June01DFS struct {
}

func (p Prac23June01DFS) numOfMinutes(n int, headId int, mgrs, informTime []int) int {
	adjList := p.createAdjList(headId, mgrs, len(informTime))
	return p.findInformTime(headId, adjList, informTime)
}

func (p Prac23June01DFS) findInformTime(empId int, adjList [][]int, informTime []int) int {
	directReports := adjList[empId]

	var maxTime int
	for _, r := range directReports {
		rTime := p.findInformTime(r, adjList, informTime)
		if maxTime < rTime {
			maxTime = rTime
		}
	}

	return maxTime + informTime[empId]
}

func (p Prac23June01DFS) createAdjList(headId int, mgrs []int, numEmps int) [][]int {
	adjList := make([][]int, numEmps, numEmps)
	for e := 0; e < len(mgrs); e++ {
		head := mgrs[e]
		if head == -1 {
			continue //ignore since head has no manager
		}

		directReports := adjList[head]
		if len(directReports) == 0 {
			directReports = make([]int, 0)
		}

		directReports = append(directReports, e)
		adjList[head] = directReports
	}

	return adjList
}
