package Q24_Course_scheduler

type Prac230601TopSort struct {
}

func (p Prac230601TopSort) canFinish(n int, preReqs [][]int) bool {
	inDegrees, adjList := p.createAdjacencyListGraph(n, preReqs)

	q := make([]int, 0)
	for c, numPreReqs := range inDegrees {
		if numPreReqs == 0 {
			q = append(q, c)
		}
	}

	var count int
	for len(q) > 0 {
		courseWithZeroPreReq := q[0]
		q = q[1:]

		count++

		deps := adjList[courseWithZeroPreReq]
		for _, c := range deps {
			inDegrees[c]--

			if inDegrees[c] == 0 {
				q = append(q, c)
			}
		}
	}

	if count == n {
		return true
	}

	return false
}

func (p Prac230601TopSort) createAdjacencyListGraph(n int, prereqs [][]int) ([]int, [][]int) {
	adjList := make([][]int, n, n)
	inDegrees := make([]int, n, n)

	for _, pr := range prereqs {
		c := pr[0]
		d := pr[1]

		inDegrees[c]++

		depCourses := adjList[d]
		if len(depCourses) == 0 {
			depCourses = make([]int, 0)
		}

		depCourses = append(depCourses, c)

		adjList[d] = depCourses
	}

	return inDegrees, adjList
}
