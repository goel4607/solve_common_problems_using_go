package course_scheduler

type Soln1BruteForce struct {
}

func (s1 Soln1BruteForce) isItPossibleToFinishAllCourses(n int, prereqs []CourseDependency) bool {
	//check for null condition
	if n <= 1 {
		return true
	}

	//create an adjacency list out of pre-requisites
	g := s1.prepareAdjacencyListGraph(n, prereqs)

	//iterate over all the n courses
	for i := 0; i < n; i++ {
		//create a new set of dependent courses
		deps := make(map[int]bool, 0)

		//traversing using bfs
		q := make([]int, 0)
		q = append(q, i) //enqueue

		for len(q) > 0 {
			d := q[0] //copy the first element
			q = q[1:] //dequeue

			if _, ok := deps[d]; ok {
				return false
			}

			if len(g[d]) == 0 {
				continue //ignore the course if it is not a pre-requisite of any course i.e. no course is dependent on it
			}
			deps[d] = true //make the pre-req
			//iterate each course's adjacency list
			for _, v := range g[d] {
				//check if this dependency is already present in seen then return false since there is a cycle among deps
				//if _, ok := deps[d]; ok {
				//	return false
				//}
				//otherwise add the dep in the dep courses set
				//deps[v] = true
				q = append(q, v)
			}
		}
	}

	//since we checked all and no cycles were found so return true
	return true
}

func (s1 Soln1BruteForce) prepareAdjacencyListGraph(n int, prereqs []CourseDependency) [][]int {
	g := make([][]int, n, n)

	for _, p := range prereqs {
		depIdx := p.preReq
		if len(g[depIdx]) == 0 {
			g[depIdx] = make([]int, 0)
		}

		//when prerequisite says [course=2,pre-req=3] then it means that 3 is prereq for 2,
		//so the adj list should populate: 2 index list with 3 in it
		g[depIdx] = append(g[depIdx], p.course)
	}

	return g
}
