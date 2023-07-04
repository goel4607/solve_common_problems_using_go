package Q20_FindNumIslands

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

type Soln1 struct {
}

const (
	WATER = iota
	LAND
)

func (s Soln1) FindNumIslands(arr [][]int) int {
	//create a boolean 2d array representing visited positions
	numRows := len(arr)
	visitedPos := make([][]bool, numRows, numRows)

	//init visitedPos arr with false values and ensure that each element is present
	numCols := len(arr[0])
	for i := range visitedPos {
		visitedPos[i] = make([]bool, numCols, numCols)
	}

	//create an int for numIslands
	var numIslands int

	//iterate from top left i.e. x=0, y=0 and iterate over each cell row by row
	for i, r := range arr {
		for j, c := range r {
			if c == WATER {
				continue //ignore if the cell is water
			}

			if visitedPos[i][j] {
				continue //ignore if the cell is already visited
			}

			numIslands++            //since first time visiting, so it's a new island, so inc numIslands
			visitedPos[i][j] = true //set the cells position to true to mark that it has been visited

			//mark all the adjoining cells in the visited places to be true by doing BFS starting from this node
			s.markAdjoiningLandAsPartOfThisIsland(Position{i, j}, arr, visitedPos)
		}
		//end of loop
	}
	//return numLands
	return numIslands
}

func (s Soln1) markAdjoiningLandAsPartOfThisIsland(p Position, arr [][]int, visitedPos [][]bool) {
	q := make([]Position, 0)
	q = append(q, p)

	for len(q) > 0 {
		e := q[0] // get the first element from Q
		q = q[1:] // remove the first element form Q

		for _, m := range MOVEMENTS {
			newPos := Position{
				r: e.r + m.r,
				c: e.c + m.c,
			}

			if s.isValidMovement(newPos, arr, visitedPos) {
				q = append(q, newPos)
				visitedPos[newPos.r][newPos.c] = true
			}
		}
	}
}

// isValidMovement checks if the position is within array bounds, is land and has not been visited before
func (s Soln1) isValidMovement(p Position, arr [][]int, visited [][]bool) bool {
	if p.r >= 0 && p.r < len(arr) {
		if p.c >= 0 && p.c < len(arr[0]) {
			if arr[p.r][p.c] == LAND && !visited[p.r][p.c] {
				return true
			}
		}
	}

	return false
}
