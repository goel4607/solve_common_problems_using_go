package dfs

type DFSSoln1 struct {
}

func (s DFSSoln1) dfs(g [][]int) []int {
	if len(g) == 0 || len(g[0]) == 0 {
		return []int{}
	}

	return s.dfsTraversal(g, make(map[int]bool), 0, make([]int, 0))
}

func (s DFSSoln1) dfsTraversal(g [][]int, seen map[int]bool, vertex int, r []int) []int {
	r = append(r, vertex)
	seen[vertex] = true

	for _, v := range g[vertex] {
		if !seen[v] {
			r = s.dfsTraversal(g, seen, v, r)
		}
	}

	return r
}
