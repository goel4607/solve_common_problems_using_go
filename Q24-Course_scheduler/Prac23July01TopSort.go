package Q24_Course_scheduler

//start time: 	2:58 pm
//end time: 	3:53 pm

type Prac23July01TopSort struct {
}

func (p Prac23July01TopSort) canFinish(numCourses int, prerequisites [][]int) bool {
	//create a adjacency list of courses where each course is an index and its dependent courses is a list of courses which are dependent upon it
	//courses 1,2,3 are dependent upon course 0, then here 0 would contain a list of its dependent courses like 0 -> [1, 2, 3]
	inDegs, adjList := p.calcPreReqsAndInDegrees(numCourses, prerequisites)

	//in top sort, we need the in degrees of each node, so we need to calculate the in degree of the course
	//say in the above example, 0 has in degree of zero and (1,2,3) has in degree of 1 each
	//so calc in-degree of nodes and adjacency list graph
	q := make([]int, 0)
	for c, inDeg := range inDegs {
		if inDeg == 0 {
			q = append(q, c)
		}
	}

	//now, we need to visit the zero in degree courses and delete this node. This will result in decreasing the
	//in-degrees of its dependent courses; similarly we need to visit all the courses such we have no course left with
	//zero in-degrees. If there are any course left with non-zero in degrees then it means we have a cycle

	var count int
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

	return count == numCourses
}

func (p Prac23July01TopSort) calcPreReqsAndInDegrees(n int, prerequisites [][]int) ([]int, [][]int) {
	adjList := make([][]int, n, n)
	inDegs := make([]int, n, n)

	for _, pr := range prerequisites {
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
