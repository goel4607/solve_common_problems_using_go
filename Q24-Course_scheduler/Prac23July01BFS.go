package Q24_Course_scheduler

//start time: 	4:59 pm
//end time:

type Prac23July01BFS struct {
}

func (p Prac23July01BFS) canFinish(numCourses int, prerequisites [][]int) bool {
	//calc adj graph
	adjList := p.calcAdjList(numCourses, prerequisites)

	//iterate over all the courses and for each course, visit its dependents and see if this course occurs in it or not
	for c := range adjList {

		//if yes, then there is a cycle, other not
		if p.isCycleExist(c, adjList) {
			return false
		}
	}

	return true
}

func (p Prac23July01BFS) isCycleExist(c int, adjList [][]int) bool {
	q := make([]int, 0)
	q = append(q, c)

	seen := make(map[int]bool)

	for len(q) > 0 {
		t := q[0]
		q = q[1:]

		if seen[t] {
			continue
		}

		seen[t] = true

		for _, d := range adjList[t] {

			if d == c {
				return true
			}

			q = append(q, d)
		}
	}

	return false
}

func (p Prac23July01BFS) calcAdjList(n int, preReqs [][]int) [][]int {
	adjList := make([][]int, n, n)
	for _, pr := range preReqs {
		c := pr[0]
		d := pr[1]

		adjs := adjList[d]
		if len(adjs) == 0 {
			adjs = make([]int, 0)
		}

		adjList[d] = append(adjs, c)
	}

	return adjList
}
