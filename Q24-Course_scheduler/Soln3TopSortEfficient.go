package Q24_Course_scheduler

type Soln3TopSortEfficient struct {
}

func (s Soln3TopSortEfficient) canFinish(n int, prereqs [][2]int) bool {
	if n <= 1 {
		return true
	}

	inDeg, g := s.createInDegreeAndAdjList(n, prereqs)

	stk := Stack{}
	for i, v := range inDeg {
		if v == 0 {
			stk.push(i)
		}
	}

	var count int
	for !stk.IsEmpty() {
		v := stk.pop()
		count++

		for _, v := range g[v] {
			inDeg[v]--
			if inDeg[v] == 0 {
				stk.push(v)
			}
		}
	}

	if count == n {
		return true
	}

	return false
}

func (s Soln3TopSortEfficient) createInDegreeAndAdjList(n int, preReqs [][2]int) ([]int, [][]int) {
	inDeg := make([]int, n, n)
	g := make([][]int, n, n)
	for _, p := range preReqs {
		if len(g[p[1]]) == 0 {
			g[p[1]] = make([]int, 0)
		}

		g[p[1]] = append(g[p[1]], p[0])

		inDeg[p[0]] += 1
	}

	return inDeg, g
}

type Stack struct {
	vals []int
}

func (stk *Stack) push(v int) {
	stk.vals = append(stk.vals, v)
}

func (stk *Stack) pop() int {
	origLen := len(stk.vals)
	v := stk.vals[origLen-1]
	stk.vals = stk.vals[:origLen-1]
	return v
}

func (stk *Stack) Len() int {
	return len(stk.vals)
}

func (stk *Stack) IsEmpty() bool {
	if stk.Len() == 0 {
		return true
	}

	return false
}
