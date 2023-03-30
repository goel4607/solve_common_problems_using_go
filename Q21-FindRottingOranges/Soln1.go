package Q21_FindRottingOranges

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

var MOVEMENTS = [4]Position{
	UP:    {-1, 0},
	RIGHT: {0, 1},
	DOWN:  {1, 0},
	LEFT:  {0, -1},
}

const (
	EMPTY = iota
	FRESH
	ROTTEN
)

type RottenPosWithOrangePos struct {
	rottenPos   []Position
	adjacentPos []Position
}

type Soln1 struct {
}

func (s Soln1) findAllOrangesRottenTime(arr [][]int) int {
	if arr == nil {
		return -1
	}

	//check num rows are > 0 if not return -1
	numRows := len(arr)
	if numRows < 1 {
		return 0
	}

	//check num cols are > 0 if not return -1
	numCols := len(arr[0])
	if numCols < 0 {
		return 0
	}

	//find all rotten orange positions and num fresh organges
	positions, numFresh := s.findRottenOrangePositions(arr)
	return s.findTotalRottingTime(positions, arr, numFresh)
}

func (s Soln1) findRottenOrangePositions(arr [][]int) ([]Position, int) {
	var positions []Position
	var fresh int

	//iterate over the array for each row and each column
	for i, r := range arr {
		for j, c := range r {
			switch c {
			case ROTTEN:
				positions = append(positions, Position{r: i, c: j})
			case FRESH:
				fresh++
			case EMPTY: //do nothing; ignore

			default:
				//need to have some check in this case!!
			}
		}
	}

	return positions, fresh
}

func (s Soln1) findTotalRottingTime(q []Position, arr [][]int, numFresh int) int {
	var ttlTime int

	currQSize := len(q)

	for len(q) > 0 {
		if currQSize == 0 {
			ttlTime++

			currQSize = len(q)
		}

		e := q[0]
		q = q[1:]

		currQSize--

		for _, m := range MOVEMENTS {
			p := Position{e.r + m.r, e.c + m.c}

			if s.isPositionValid(p, arr) {
				arr[p.r][p.c] = ROTTEN
				numFresh--

				q = append(q, p)
			}
		}
	}

	if numFresh > 0 {
		return -1
	}

	return ttlTime
}

func (s Soln1) isPositionValid(p Position, arr [][]int) bool {
	if p.r >= 0 && p.r < len(arr) {
		if p.c >= 0 && p.c < len(arr[0]) {
			if arr[p.r][p.c] == FRESH {
				return true
			}
		}
	}

	return false
}
