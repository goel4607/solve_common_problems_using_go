package bfs

type Soln1 struct {
}

type Position struct {
	r int
	c int
}

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

var DIRECTIONS = [4]Position{
	UP:    {-1, 0},
	RIGHT: {0, 1},
	DOWN:  {1, 0},
	LEFT:  {0, -1},
}

func (s Soln1) BFS(arr [][]int, r, c int) []int {
	visited := make([][]bool, len(arr), len(arr))
	for i := range visited {
		numCols := len(arr[i])
		visited[i] = make([]bool, numCols, numCols)
	}

	q := make([]Position, 0, len(arr)*len(arr[0]))
	q = append(q, Position{r, c}) // add this position in the Q

	visited[r][c] = true

	result := make([]int, 0, len(arr)*len(arr[0]))
	result = append(result, arr[r][c])

	return s.visitBFS(arr, q, visited, result)
}

func (s Soln1) visitBFS(arr [][]int, q []Position, visited [][]bool, result []int) []int {
	for len(q) > 0 {
		p := q[0] // get the first element
		q = q[1:] // remove the first element from q

		for _, d := range DIRECTIONS {
			newP := s.computeValidPosition(p, d, visited)
			if newP != nil && !visited[newP.r][newP.c] {
				visited[newP.r][newP.c] = true
				result = append(result, arr[newP.r][newP.c])
				q = append(q, *newP)
			}
		}
	}

	return result
}

func (s Soln1) computeValidPosition(p Position, d Position, visited [][]bool) *Position {
	r := p.r + d.r
	if r >= 0 && r < len(visited) {
		c := p.c + d.c
		if c >= 0 && c < len(visited[0]) {
			return &Position{r, c}
		}
	}

	return nil
}
