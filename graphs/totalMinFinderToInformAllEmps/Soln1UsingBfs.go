package totalMinFinderToInformAllEmps

type Soln1UsingBfs struct {
}

//goland:noinspection SpellCheckingInspection
func (s Soln1UsingBfs) findTotalMinutesToInformAllEmps(headId int, mgrs, informTime []int) int {
	if len(mgrs) <= 1 {
		return 0
	}
	//make adj list say g for the mgrs
	g := s.makeAdjListGraph(mgrs, headId)
	//traverse g with bfs starting at the headID

	q := make([]int, 0)
	q = append(q, headId)

	totalTime := informTime[headId]

	for len(q) > 0 {
		mgr := q[0]
		q = q[1:]

		reportees := g[mgr]

		//while doing bfs track the timings for each parent and consider the max and add the max after all child's evaluations
		var maxTime int //calc the max time it takes for each mgr to inform its reportee
		for _, empId := range reportees {
			if informTime[empId] == 0 {
				//0 inform time indicates that it is not a manager
				//ignore this entry since there are no further employees to inform under this employee
				continue
			}

			if maxTime < informTime[empId] {
				maxTime = informTime[empId]
			}

			q = append(q, empId) //add all the manager reportees since we need to calc each one's inform time separately
		}

		totalTime += maxTime
	}

	return totalTime
}

func (s Soln1UsingBfs) makeAdjListGraph(mgrs []int, headId int) [][]int {
	g := make([][]int, len(mgrs), len(mgrs))

	for i, v := range mgrs {
		if i == headId {
			continue //ignore since head does not have any manager
		}

		if len(g[v]) == 0 {
			g[v] = make([]int, 0)
		}

		g[v] = append(g[v], i)
	}

	return g
}
