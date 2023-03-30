package Q23_TotalMinFinderToInformAllEmps

type Prac23Week11UsingDFS struct {
}

func (p Prac23Week11UsingDFS) findTotalMinutesToInformAllEmps(headId int, mgrs, informTime []int) int {
	if len(mgrs) <= 1 {
		return 0
	}
	//make adjacency list of mgrs array such each manager lists its employee id's
	g := p.createAdjacencyListWithMgrToReportees(mgrs)
	total := informTime[headId]
	return total + p.findMinTime(headId, g, informTime)
}

func (p Prac23Week11UsingDFS) findMinTime(m int, g [][]int, informTime []int) int {
	var max int
	for _, v := range g[m] {
		if informTime[v] == 0 {
			continue
		}

		empHierMax := p.findMinTime(v, g, informTime)
		empMax := empHierMax + informTime[v]
		if max < empMax {
			max = empMax
		}
	}

	return max
}

func (p Prac23Week11UsingDFS) createAdjacencyListWithMgrToReportees(mgrs []int) [][]int {
	const HEAD_MGR = -1

	g := make([][]int, len(mgrs), len(mgrs))
	for i, v := range mgrs {
		if v == HEAD_MGR {
			continue
		}
		emps := g[v]
		if len(emps) == 0 {
			emps = make([]int, 0)
		}

		emps = append(emps, i)
		g[v] = emps
	}

	return g
}
