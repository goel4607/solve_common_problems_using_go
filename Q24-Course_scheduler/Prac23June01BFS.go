package Q24_Course_scheduler

type Prac230601BFS struct {
}

func (p Prac230601BFS) canFinish(n int, prereqs [][2]int) bool {
	//create adjList such that prereq (arr index) contains the list of its dependent courses
	adjList := p.createAdjacencyList(n, prereqs)

	//iterate over the graph and for each vertex if it has creates a cycle with any of its dependency graph
	for i, depCourses := range adjList {

		if len(depCourses) == 0 {
			continue
		}

		if p.checkDepCyclePresent(i, adjList) {
			return false
		}
	}

	return true
}

// we want to iterate over each pre-req using BFS to find out if any course is again referenced by it
func (p Prac230601BFS) checkDepCyclePresent(startingCourse int, adjList [][]int) bool {
	seenCourses := make(map[int]bool)

	q := make([]int, 0)

	mainDeps := adjList[startingCourse]
	for _, d := range mainDeps {
		q = append(q, d)
	}

	for len(q) > 0 {
		currCourse := q[0]
		q = q[1:]

		if seenCourses[startingCourse] {
			return true
		}

		if seenCourses[currCourse] { //optimization
			continue
		}

		seenCourses[currCourse] = true
		//check each dependent course to see if it has already being visited; if yes, then we found the cycle;
		//otherwise keep checking
		for _, depCourseId := range adjList[currCourse] {
			q = append(q, depCourseId)
		}
	}

	return false
}

func (p Prac230601BFS) createAdjacencyList(n int, preReqs [][2]int) [][]int {
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
