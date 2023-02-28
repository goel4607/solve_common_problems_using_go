package dfs

type Soln1 struct {
}

type Position struct {
	row int
	col int
}

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

var DIRECTIONS = []Position{
	UP:    {-1, 0},
	RIGHT: {0, 1},
	DOWN:  {+1, 0},
	LEFT:  {0, -1},
}

func (s Soln1) DFS(arr [][]int) []int {
	visitedPos := make([][]bool, len(arr), len(arr))
	for i := range visitedPos {
		visitedPos[i] = make([]bool, len(arr[i]), len(arr[i]))
	}

	r := make([]int, 0, len(arr)*len(arr))
	r = append(r, arr[0][0])
	visitedPos[0][0] = true
	return s.visitDFS(arr, Position{0, 0}, visitedPos, r)
}

func (s Soln1) visitDFS(arr [][]int, current Position, visitedPos [][]bool, r []int) []int {
	for _, d := range DIRECTIONS {
		p := s.getEventualPosition(arr, current, d)
		if p != nil && !visitedPos[p.row][p.col] {
			visitedPos[p.row][p.col] = true
			r = append(r, arr[p.row][p.col])
			r = s.visitDFS(arr, *p, visitedPos, r)
		}
	}

	return r
}

func (s Soln1) getEventualPosition(arr [][]int, current, d Position) *Position {
	r := current.row + d.row
	if r >= 0 && r < len(arr) {
		c := current.col + d.col
		if c >= 0 && c < len(arr[r]) {
			return &Position{r, c}
		}
		return nil
	}

	return nil
}
