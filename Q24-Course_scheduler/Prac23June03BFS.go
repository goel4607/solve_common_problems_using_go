package Q24_Course_scheduler

// Start time: 	11:11 pm
// End Time:	11:23 pm

type Prac23June03BFS struct {
}

func (p Prac23June03BFS) canFinish(n int, preReqs [][]int) bool {
	//create an adjacency list
	adjList := p.createAdjList(n, preReqs)

	//we can iterate over courses in adj list and find out if that course is revisited during its BFS,
	for i := range adjList {
		if p.checkIfCycleExists(i, adjList) {
			return false
		}
	}

	//if yes then there is a cycle otherwise not
	return true
}

func (p Prac23June03BFS) checkIfCycleExists(c int, adjList [][]int) bool {
	q := make([]int, 0)
	q = append(q, c)

	for len(q) > 0 {
		v := q[0]
		q = q[1:]

		for _, d := range adjList[v] {
			if d == c {
				return true
			}

			q = append(q, d)
		}
	}

	return false
}

func (p Prac23June03BFS) createAdjList(n int, preReqs [][]int) [][]int {
	adjList := make([][]int, n, n)

	for _, pr := range preReqs {
		c := pr[0]
		d := pr[1]

		deps := adjList[d]
		if len(deps) == 0 {
			deps = make([]int, 0)
		}

		adjList[d] = append(deps, c)
	}

	return adjList
}
