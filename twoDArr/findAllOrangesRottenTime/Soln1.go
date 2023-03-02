package findAllOrangesRottenTime

import "errors"

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
	NO_ORANGE = iota
	ORANGE
	ROTTEN_ORANGE
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
		return -1
	}

	//check num cols are > 0 if not return -1
	numCols := len(arr[0])
	if numCols < 0 {
		return -1
	}

	//find num islands with rotten orange position in it
	allRpos, err := s.findNumIslandsWithRottenOrangePositions(arr)
	if err != nil {
		return -1 //check if there is any island where there is no rotten orange then return -1
	}

	visited := make([][]bool, len(arr), len(arr))
	for i := range visited {
		visited[i] = make([]bool, len(arr[0]), len(arr[0]))
	}

	//iterate over each rotten orange on each island and compute total time when all oranges will be rotten
	var ttlMins int
	for _, r := range allRpos {
		ttlMins += s.findAllRottenOrangeTime(r, arr, visited)
	}

	return ttlMins
}

func (s Soln1) findNumIslandsWithRottenOrangePositions(arr [][]int) ([]*RottenPosWithOrangePos, error) {
	//a boolean arr indicating what positions have been visited
	visited := make([][]bool, len(arr), len(arr))
	for i := range visited {
		visited[i] = make([]bool, len(arr[0]), len(arr[0]))
	}

	allRpos := make([]*RottenPosWithOrangePos, 0)

	//iterate over the array for each row and each column
	for i, r := range arr {
		for j, c := range r {
			if c == NO_ORANGE { //ignore if the element is not orange
				continue
			}

			if visited[i][j] { //ignore if it has been visited before
				continue
			}

			//findAdjoiningElements along with rotten orange positions
			rpos := s.findAdjoiningElements(Position{i, j}, arr, visited)
			//if num rotten orange positions is zero and num adjoining pos is non-zero, then return -1 indicating that
			if len(rpos.rottenPos) < 1 && len(rpos.adjacentPos) > 0 {
				return nil, errors.New("an orange bunch exists with no rotten oranges")
			}

			allRpos = append(allRpos, rpos)
		}
	}

	return allRpos, nil
}

func (s Soln1) findAdjoiningElements(p Position, arr [][]int, visited [][]bool) *RottenPosWithOrangePos {
	roPos := make([]Position, 0)
	oPos := make([]Position, 0)

	q := make([]Position, 0)
	q = append(q, p)

	for len(q) > 0 {
		e := q[0]
		q = q[1:]

		visited[e.r][e.c] = true

		switch arr[e.r][e.c] {
		case ROTTEN_ORANGE:
			roPos = append(roPos, e)
		case ORANGE:
			oPos = append(oPos, e)
		}

		for _, m := range MOVEMENTS {
			newPos := Position{e.r + m.r, e.c + m.c}
			if s.isPositionValid(newPos, arr, visited) {
				q = append(q, newPos)
			}
		}
	}

	return &RottenPosWithOrangePos{roPos, oPos}
}

func (s Soln1) isPositionValid(p Position, arr [][]int, visited [][]bool) bool {
	if p.r >= 0 && p.r < len(arr) {
		if p.c >= 0 && p.c < len(arr[0]) {
			if arr[p.r][p.c] != NO_ORANGE && !visited[p.r][p.c] {
				return true
			}
		}
	}

	return false
}

func (s Soln1) findAllRottenOrangeTime(rpos *RottenPosWithOrangePos, arr [][]int, visited [][]bool) int {

	var ttlMins int

	//starting from the rotten position iterate over the array for the
	q := make([]Position, 0)
	q = append(q, rpos.rottenPos...)

	for len(q) > 0 {
		e := q[0]
		q = q[1:]

		visited[e.r][e.c] = true

		var numPos int
		for _, m := range MOVEMENTS {
			newPos := Position{e.r + m.r, e.c + m.c}
			if s.isPositionValid(newPos, arr, visited) {
				q = append(q, newPos)
				numPos++
			}
		}

		if numPos > 0 {
			ttlMins++
		}
	}

	return ttlMins
}
