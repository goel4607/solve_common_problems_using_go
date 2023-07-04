package Q24_Course_scheduler

type Prac23June04TopSort struct {
}

func (p Prac23June04TopSort) canFinish(numCourses int, prerequisites [][]int) bool {
	inDegs, adjs := p.createAdjList(numCourses, prerequisites)

	q := make([]int, 0)
	for i := 0; i < len(inDegs); i++ {
		if inDegs[i] == 0 {
			q = append(q, i)
		}
	}

	var count int
	for len(q) > 0 {
		c := q[0]
		q = q[1:]

		count++

		for _, c := range adjs[c] {
			inDegs[c]--
			if inDegs[c] == 0 {
				q = append(q, c)
			}
		}
	}

	if count == numCourses {
		return true
	}

	return false
}

func (p Prac23June04TopSort) createAdjList(n int, preReqs [][]int) ([]int, [][]int) {
	adjList := make([][]int, n, n)
	inDegs := make([]int, n)

	for _, pr := range preReqs {
		c := pr[0]
		d := pr[1]

		inDegs[c]++

		adjs := adjList[d]
		if len(adjs) == 0 {
			adjs = make([]int, 0)
		}

		adjList[d] = append(adjs, c)
	}

	return inDegs, adjList
}
