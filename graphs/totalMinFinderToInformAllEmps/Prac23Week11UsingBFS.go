package totalMinFinderToInformAllEmps

type Prac23Week11UsingBfs struct {
}

func (p Prac23Week11UsingBfs) findTotalMinutesToInformAllEmps(headId int, mgrs, informTime []int) int {
	if len(mgrs) <= 1 {
		return 0
	}
	//make adjacency list of mgrs array such each manager lists its employee id's
	g := p.createAdjacencyListWithMgrToReportees(mgrs)
	//declare total time say total
	var total int
	total += informTime[headId]
	//make a q and add the headId into it
	q := make([]int, 0)
	q = append(q, headId)

	//iterate over the Q till is not empty
	for len(q) > 0 {
		//take an emp from Q
		m := q[0]
		q = q[1:]

		//iterate over its adjacent elements and find the max out of them and keep adding then to the queue
		var max int
		for _, empId := range g[m] {
			if informTime[empId] == 0 {
				continue
			}

			if max < informTime[empId] {
				max = informTime[empId]
			}

			q = append(q, empId)
		}

		//inc total with max value
		total += max
	}
	//return total
	return total
}

func (p Prac23Week11UsingBfs) createAdjacencyListWithMgrToReportees(mgrs []int) [][]int {
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
