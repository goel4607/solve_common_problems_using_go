package Q24_Course_scheduler

// Start time: 	10:45 pm
// End Time:	11:10 pm

type Prac23June03TopSort struct {
}

func (p Prac23June03TopSort) canFinish(numCourses int, preReqs [][]int) bool {
	//create a adjList map and in-degrees arr
	adjList, inDegs := p.createAdjListAndInDegrees(numCourses, preReqs)

	//iterate over adjList map and find courses with zero in-degrees
	//add those courses into a Q
	q := make([]int, 0)
	for i := range inDegs {
		if inDegs[i] == 0 {
			q = append(q, i)
		}
	}

	var count int
	//iterate over the Q while it is not empty and inc count for every element
	for len(q) > 0 {
		c := q[0]
		q = q[1:]

		count++

		for _, d := range adjList[c] {
			inDegs[d]--
			if inDegs[d] == 0 {
				q = append(q, d)
			}
		}
	}

	//cycle does not exist if the count matches the numCourses, return true
	if count == numCourses {
		return true
	}

	//otherwise cycle exists, return false
	return false
}

func (p Prac23June03TopSort) createAdjListAndInDegrees(numCourses int, preReqs [][]int) ([][]int, []int) {
	adjList := make([][]int, numCourses, numCourses)
	inDegs := make([]int, numCourses, numCourses)

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

	return adjList, inDegs
}
