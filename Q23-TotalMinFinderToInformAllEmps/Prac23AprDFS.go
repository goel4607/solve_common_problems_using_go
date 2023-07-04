package Q23_TotalMinFinderToInformAllEmps

// start time: 9:20 pm
type Prac23AprDFS struct {
}

func (s Prac23AprDFS) numOfMinutes(headId int, manager, informTime []int) int {
	g := s.makeAdjList(headId, manager)
	return s.visitEachSubOrdinateManagerUsingDFS(headId, g, informTime)
}

// makeAdjList makes an adjacency list such that each index represents a manager and lists all the employees under it
// emp is rep as index
// value at mgrs[0] is the manager of employee 0
// value at mgrs[1] is the manager of employee 1
func (s Prac23AprDFS) makeAdjList(headId int, mgrs []int) [][]int {
	adjList := make([][]int, len(mgrs), len(mgrs))
	for e, m := range mgrs {
		if e == headId {
			continue // head does not have any manager so ignore this entry
		}
		if adjList[m] == nil {
			emps := make([]int, 0)
			emps = append(emps, e)
			adjList[m] = emps
		} else {
			adjList[m] = append(adjList[m], e)
		}
	}

	return adjList
}

func (s Prac23AprDFS) visitEachSubOrdinateManagerUsingDFS(mgrId int, g [][]int, informTime []int) int {
	emps := g[mgrId]
	var maxTime int
	for _, e := range emps {
		mgrTime := s.visitEachSubOrdinateManagerUsingDFS(e, g, informTime)
		if maxTime < mgrTime {
			maxTime = mgrTime
		}
	}

	return maxTime + informTime[mgrId]
}
