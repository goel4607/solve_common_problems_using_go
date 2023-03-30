package Q22_WallsAndGates

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

var Movements = []Position{
	UP:    {-1, 0},
	RIGHT: {0, 1},
	DOWN:  {1, 0},
	LEFT:  {0, -1},
}

type Soln1 struct {
}

func (s Soln1) FillRoomsWithNumHopsToNearestGate(arr [][]int) {
	if len(arr) < 1 {
		return
	}

	if len(arr[0]) < 1 {
		return
	}

	//find the gate positions by iterating over the arr
	gates := s.findAllGatePositions(arr)

	//iterate over each gate
	for _, gPos := range gates {
		s.traverseDfsAndMarkRoomsWithDistance(arr, gPos, 1)
	}
}

func (s Soln1) findAllGatePositions(arr [][]int) []Position {
	var gates []Position
	for i, r := range arr {
		for j, c := range r {
			if c == GATE {
				gates = append(gates, Position{i, j})
			}
		}
	}

	return gates
}

// traverse the arr starting from each gate position using dfs
// while traversing fill each position incrementing from zero and next to the current position with +1
// when filling set the value to min between its value and possible current position
// after this traversal all the rooms will be filled wit the distance from this gate
func (s Soln1) traverseDfsAndMarkRoomsWithDistance(arr [][]int, now Position, dist int) {

	for _, p := range Movements {
		move := Position{p.r + now.r, p.c + now.c}
		if s.isMovementValid(arr, move) {
			if dist < arr[move.r][move.c] {
				arr[move.r][move.c] = dist //update the min value
				s.traverseDfsAndMarkRoomsWithDistance(arr, move, dist+1)
			}
		}
	}
}

func (s Soln1) isMovementValid(arr [][]int, p Position) bool {
	if p.r >= 0 && p.r < len(arr) {
		if p.c >= 0 && p.c < len(arr[0]) {
			if arr[p.r][p.c] != WALL && arr[p.r][p.c] != GATE {
				return true
			}
		}
	}

	return false
}
