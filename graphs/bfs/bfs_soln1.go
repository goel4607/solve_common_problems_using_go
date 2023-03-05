package bfs

type BFSSoln1 struct {
}

func (s BFSSoln1) bfs(g [][]int) []int {
	if len(g) == 0 || len(g[0]) == 0 {
		return []int{}
	}

	seen := make(map[int]bool)

	q := make([]int, 0)
	q = append(q, 0)

	r := make([]int, 0)

	for len(q) > 0 {
		vertex := q[0]
		q = q[1:]

		if !seen[vertex] {
			r = append(r, vertex)
			seen[vertex] = true
		}

		adj := g[vertex]
		for _, v := range adj {
			if !seen[v] {
				q = append(q, v)
			}
		}
	}

	return r
}
