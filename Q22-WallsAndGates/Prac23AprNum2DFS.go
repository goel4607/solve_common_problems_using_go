package Q22_WallsAndGates

type Pos2 struct {
	r int
	c int
}

const (
	Up2 = iota
	Right2
	Down2
	Left2
)

var Moves2 = []Pos2{
	Up2:    {-1, 0},
	Right2: {0, 1},
	Down2:  {1, 0},
	Left2:  {0, -1},
}

type Prac23AprNum2DFS struct {
}

func (s Prac23AprNum2DFS) FillRoomsWithNumHopsToNearestGate(arr [][]int) {
	var gates []Pos2
	//iterate over the arr and find out the gates
	for i, row := range arr {
		for j := range row {
			if arr[i][j] == GATE {
				gates = append(gates, Pos2{i, j})
			}
		}
	}

	for _, p := range gates {
		s.updateCellStepsUsingDFS(arr, p, 1)
	}
}

// when a gate is found then using BFS visit all the possible neighbours and update its distance from this gate
// if it is lower that the one specified
func (s Prac23AprNum2DFS) updateCellStepsUsingDFS(arr [][]int, p Pos2, steps int) {
	for _, m := range Moves2 {
		if s.isValidMove(arr, p, m) {
			t := Pos2{r: p.r + m.r, c: p.c + m.c}
			if arr[t.r][t.c] > steps {
				arr[t.r][t.c] = steps

				s.updateCellStepsUsingDFS(arr, t, steps+1)
			}
		}
	}
}

func (s Prac23AprNum2DFS) isValidMove(arr [][]int, v, m Pos2) bool {
	r := v.r + m.r
	if r < 0 || r >= len(arr) {
		return false
	}

	c := v.c + m.c
	if c < 0 || c >= len(arr[0]) {
		return false
	}

	if arr[r][c] == WALL {
		return false
	}

	if arr[r][c] == GATE {
		return false
	}

	return true
}
