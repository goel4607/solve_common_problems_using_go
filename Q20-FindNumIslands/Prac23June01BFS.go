package Q20_FindNumIslands

// start time: 	8:35 pm
// end time:	9:13 pm
// ---------------------
// time taken: 	38 minutes

type Prac23June01BFS struct {
}

func (p Prac23June01BFS) FindNumIslands(arr [][]byte) int {
	//create a 2d arr of same size of boolean
	numRows, numCols := len(arr), len(arr[0])
	if numRows == 0 || numCols == 0 {
		return 0
	}
	
	var numIslands int
	//iterate over the arr from the top left and call islandMarker func
	for i, row := range arr {
		for j := range row {
			//if a non-seen LAND is found then markIsland func will be called and it will mark it again
			if arr[i][j] == 1 {
				p.markIsland(Point{i, j}, arr)

				//inc the num islands
				numIslands++
			}
		}
	}

	//so after iterating all the elements, the number of islands is found
	return numIslands
}

type Point struct {
	r int
	c int
}

var MOVES = [4]Point{
	0: {-1, 0},
	1: {0, 1},
	2: {1, 0},
	3: {0, -1},
}

// this func will use BFS and mark all the places where LAND is present in seen array and return when no element is remaining
func (p Prac23June01BFS) markIsland(pos Point, arr [][]byte) {
	q := make([]*Point, 0)
	q = append(q, &pos)

	for len(q) > 0 {
		currPos := q[0]
		q = q[1:]

		for _, m := range MOVES {
			if vPoint, ok := p.isValidMove(*currPos, m, arr); ok {
				arr[vPoint.r][vPoint.c] = 0
				q = append(q, vPoint)
			}
		}
	}
}

func (p Prac23June01BFS) isValidMove(pos, m Point, arr [][]byte) (*Point, bool) {
	r := pos.r + m.r
	if r < 0 || r >= len(arr) {
		return nil, false
	}

	c := pos.c + m.c
	if c < 0 || c >= len(arr[0]) {
		return nil, false
	}

	if arr[r][c] == 0 {
		return nil, false
	}

	if arr[r][c] == 1 {
		return &Point{r, c}, true
	}

	return nil, false
}
