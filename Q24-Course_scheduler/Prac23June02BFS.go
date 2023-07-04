package Q24_Course_scheduler

// time: 10:41 = 11
type Prac230602BFS struct {
}

func (p Prac230602BFS) canFinish(n int, preReqs [][2]int) bool {
	//create adjacency list
	adjList := p.createAdjacencyListGraph(n, preReqs)
	//iterate over each course, do bfs, with keeping a track of course dependencies, see if any of the dependencies
	for c := range adjList {
		//are present in the tracked deps, if yes, then we found a cycle and return false
		if p.checkDependencyCycle(c, adjList) {
			return false
		}
	}

	return true
}

func (p Prac230602BFS) checkDependencyCycle(startingCourse int, adjList [][]int) bool {
	seenCourses := make(map[int]bool)

	q := make([]int, 0)

	mainDeps := adjList[startingCourse]
	for _, d := range mainDeps {
		q = append(q, d)
	}

	for len(q) > 0 {
		currCourse := q[0]
		q = q[1:]

		if seenCourses[startingCourse] { //found the starting node so found a cycle
			return true // found a cycle
		}

		if seenCourses[currCourse] { // optimization
			continue
		}

		seenCourses[currCourse] = true

		//visit each of the deps and see if anyone has been seen earlier, if yes, then there is a cycle
		for _, d := range adjList[currCourse] {
			q = append(q, d)
		}
	}

	return false
}

func (p Prac230602BFS) createAdjacencyListGraph(n int, preReqs [][2]int) [][]int {
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
