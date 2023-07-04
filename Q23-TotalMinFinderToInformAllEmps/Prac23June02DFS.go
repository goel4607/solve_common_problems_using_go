package Q23_TotalMinFinderToInformAllEmps

// Start time:	2:35 pm
// End time:	2:50 pm

type Prac23June02DFS struct {
}

func (p Prac23June02DFS) numOfMinutes(n int, headId int, manager, informTime []int) int {
	if len(manager) == 0 {
		return 0
	}
	adjList := p.prepareAdjList(manager)
	return p.findMaxManagerInformTime(headId, informTime, adjList)
}

func (p Prac23June02DFS) findMaxManagerInformTime(m int, informTime []int, adjList [][]int) int {
	var max int
	for _, e := range adjList[m] {
		t := p.findMaxManagerInformTime(e, informTime, adjList)
		if max < t {
			max = t
		}
	}

	return informTime[m] + max
}

func (p Prac23June02DFS) prepareAdjList(manager []int) [][]int {
	adjList := make([][]int, len(manager), len(manager))

	for e, m := range manager {
		if m == -1 {
			continue // head does not have any manager
		}
		emps := adjList[m]
		if len(emps) == 0 {
			emps = make([]int, 0)
		}

		adjList[m] = append(emps, e)
	}

	return adjList
}
