package Q22_WallsAndGates

type Pos1 struct {
	r int
	c int
}

const (
	Up1 = iota
	Right1
	Down1
	Left1
)

var Moves = []Pos1{
	Up1:    {-1, 0},
	Right1: {0, 1},
	Down1:  {1, 0},
	Left1:  {0, -1},
}

type Prac23AprNum1BFS struct {
}

func (s Prac23AprNum1BFS) FillRoomsWithNumHopsToNearestGate(arr [][]int) {
	var gates []Pos1
	//iterate over the arr and find out the gates
	for i, row := range arr {
		for j := range row {
			if arr[i][j] == GATE {
				gates = append(gates, Pos1{i, j})
			}
		}
	}

	for _, p := range gates {
		s.updateCellStepsUsingBFS(arr, p)
	}
}

// when a gate is found then using BFS visit all the possible neighbours and update its distance from this gate
// if it is lower that the one specified
func (s Prac23AprNum1BFS) updateCellStepsUsingBFS(arr [][]int, p Pos1) {
	var q []Pos1
	q = append(q, p)

	var stepsQ []int
	stepsQ = append(stepsQ, 0)

	for len(q) > 0 {
		v := q[0]
		q = q[1:]

		steps := stepsQ[0]
		stepsQ = stepsQ[1:]

		for _, m := range Moves {
			if s.isValidMove(arr, v, m) {
				t := Pos1{r: v.r + m.r, c: v.c + m.c}
				if arr[t.r][t.c] > steps+1 {
					arr[t.r][t.c] = steps + 1
					q = append(q, t)

					stepsQ = append(stepsQ, steps+1)
				}
			}
		}
	}
}

func (s Prac23AprNum1BFS) isValidMove(arr [][]int, v, m Pos1) bool {
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

type Stack []int

func (st Stack) Push(p int) {
	st = append(st, p)
}

func (st Stack) Pop() int {
	oldLen := len(st)
	v := st[oldLen-1]
	st = st[:oldLen-1]

	return v
}

func (st Stack) IsEmpty() bool {
	if len(st) == 0 {
		return true
	}

	return false
}
