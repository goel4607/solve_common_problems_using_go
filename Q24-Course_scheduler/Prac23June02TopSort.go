package Q24_Course_scheduler

//Time: 12:35 -> 12:52

type Prac23June02TopSort struct {
}

func (p Prac23June02TopSort) canFinish(numCourses int, prerequisites [][]int) bool {
	//find the in degrees of all the courses && create a adj list of dep courses
	inDegs, adjList := p.createInDegAndAdjList(numCourses, prerequisites)
	//find courses with zero in-degrees
	q := make([]int, 0)
	for c, inDeg := range inDegs {
		if inDeg == 0 {
			q = append(q, c)
		}
	}

	//keep iterating over the zero in-degrees and simulate removal these course deps by decreasing the in-degrees of
	//each of the dependent courses;
	var count int
	for len(q) > 0 {
		c := q[0]
		q = q[1:]

		count++

		//keep checking if the in-degree is zero if yes, then keep iterating over it
		deps := adjList[c]
		for _, d := range deps {
			inDegs[d]--

			if inDegs[d] == 0 {
				q = append(q, d)
			}
		}
	}

	//check if the count of courses with zero in-degrees is same as the number of courses, if yes, then we do not have
	//cycle
	return count == numCourses
}

func (p Prac23June02TopSort) createInDegAndAdjList(n int, preReqs [][]int) ([]int, [][]int) {
	//find the in degrees of all the courses
	inDegs := make([]int, n, n)
	//create a adj list of dep courses
	adjList := make([][]int, n, n)

	for _, pr := range preReqs {
		c := pr[0]
		d := pr[1]

		inDegs[c]++

		deps := adjList[d]
		if len(deps) == 0 {
			deps = make([]int, 0)
		}

		adjList[d] = append(deps, c)
	}

	return inDegs, adjList
}
