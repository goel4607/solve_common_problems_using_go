package Q22_WallsAndGates

// start time: 	12:38 pm
// end time:	 1:02 pm
// ---------------------
// time taken: 	 44 minutes

type Point struct {
	r int
	c int
}

type Prac23June01DFS struct {
}

func (p Prac23June01DFS) FillRoomsWithNumHopsToNearestGate(arr [][]int) {
	gates := p.findGatePositions(arr)
	for _, g := range gates {
		p.visitEmptyCells(g, 0, arr)
	}
}

var June01Move = [4]Point{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func (p Prac23June01DFS) visitEmptyCells(curr Point, currNum int, arr [][]int) {
	for _, m := range June01Move {
		if vPoint, ok := p.isValidMove(curr, m, arr, currNum+1); ok {
			arr[vPoint.r][vPoint.c] = currNum + 1
			p.visitEmptyCells(*vPoint, currNum+1, arr)
		}
	}
}

func (p Prac23June01DFS) isValidMove(pos, mov Point, arr [][]int, currNum int) (*Point, bool) {
	r := pos.r + mov.r
	if r < 0 || r >= len(arr) {
		return nil, false
	}

	c := pos.c + mov.c
	if c < 0 || c >= len(arr[0]) {
		return nil, false
	}

	if arr[r][c] == WALL || arr[r][c] == GATE {
		return nil, false
	}

	if arr[r][c] <= currNum {
		return nil, false
	}

	return &Point{r, c}, true
}

func (p Prac23June01DFS) findGatePositions(arr [][]int) []Point {
	gates := make([]Point, 0)
	for r, row := range arr {
		for c := range row {
			if arr[r][c] == GATE {
				gates = append(gates, Point{r, c})
			}
		}
	}

	return gates
}
