package Q24_Course_scheduler

// start time: 10.57 pm
type Prac23AprTopSort struct {
}

func (s Prac23AprTopSort) canFinish(n int, prereqs [][2]int) bool {
	//create in-degree list and adjacency list for each course's next courses
	inDegs, adjs := s.createInDegreeListAndAdjList(n, prereqs)
	//find zero in-degrees to start with
	stack := MyStack{}
	for i, d := range inDegs {
		if d == 0 {
			stack.Push(i)
		}
	}

	//iterate over adjacent courses from zero indegree
	var count int
	for len(stack) > 0 {
		zeroInDeg := stack.Pop()

		//keep a count of zero in-deg courses
		count++

		//decrease each course's in-degree
		for _, v := range adjs[zeroInDeg] {
			inDegs[v]--

			//during decreasing in-degrees add the courses with zero in-degrees into the list
			if inDegs[v] == 0 {
				stack.Push(v)
			}
		}
	}

	//check if the num zero in-deg courses is same as num courses then there is no cycles, otherwise there is a cycle
	if count == n {
		return true
	}

	return false
}

func (s Prac23AprTopSort) createInDegreeListAndAdjList(n int, prereqs [][2]int) ([]int, [][]int) {
	inDegs := make([]int, n, n)
	adjs := make([][]int, n, n)
	for _, pr := range prereqs {
		inDegs[pr[0]] += 1

		if adjs[pr[1]] == nil {
			a := make([]int, 0)
			a = append(a, pr[0])

			adjs[pr[1]] = a
		} else {
			adjs[pr[1]] = append(adjs[pr[1]], pr[0])
		}
	}

	return inDegs, adjs
}

type MyStack []int

func (s *MyStack) Push(v int) {
	*s = append(*s, v)
}

func (s *MyStack) Pop() int {
	oldLen := len(*s)
	e := (*s)[oldLen-1]
	*s = (*s)[0 : oldLen-1]
	return e
}
