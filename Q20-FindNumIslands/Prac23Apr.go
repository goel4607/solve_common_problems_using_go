package Q20_FindNumIslands

type Pos1 struct {
	r int
	c int
}

const (
	Up = iota
	Down
	Left
	Right
)

var BFSMoves = []Pos1{
	Up:    {-1, 0},
	Down:  {1, 0},
	Left:  {0, -1},
	Right: {0, 1},
}

type Prac23Apr struct {
}

func (s Prac23Apr) findNumIslands(arr [][]int) int {
	numRows := len(arr)
	if numRows == 0 {
		return 0
	}

	numCols := len(arr[0])
	if numCols == 0 {
		return 0
	}

	visited := make([][]bool, numRows, numRows)
	for i := range visited {
		visited[i] = make([]bool, numCols, numCols)
	}

	var islands int

	for i, row := range arr {
		for j := range row {
			if arr[i][j] == WATER {
				continue
			}

			if visited[i][j] {
				continue //ignore as it is already visited
			}
			s.visitBFS(arr, Pos1{r: i, c: j}, visited)
			islands++
		}
	}

	return islands
}

func (s Prac23Apr) visitBFS(arr [][]int, p Pos1, visited [][]bool) {
	MaxRows, MaxCols := len(arr), len(arr[0])

	var q []Pos1
	q = append(q, p)

	for len(q) > 0 {
		v := q[0]
		q = q[1:]

		visited[v.r][v.c] = true

		for _, m := range BFSMoves {
			if s.isValidMove(arr, visited, v, m, MaxRows, MaxCols) {
				bfsPos := Pos1{v.r + m.r, v.c + m.c}
				q = append(q, bfsPos)
			}
		}
	}
}

func (s Prac23Apr) isValidMove(arr [][]int, visited [][]bool, p Pos1, m Pos1, MaxRows, MaxCols int) bool {
	r := p.r + m.r
	if r < 0 || r >= MaxRows {
		return false
	}

	c := p.c + m.c
	if c < 0 || c >= MaxCols {
		return false
	}

	if visited[r][c] {
		return false //ignore if already visited
	}

	if arr[r][c] == WATER {
		return false
	}

	return true
}
