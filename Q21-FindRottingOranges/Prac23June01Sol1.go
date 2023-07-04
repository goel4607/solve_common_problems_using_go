package Q21_FindRottingOranges

// start time: 	11:00 pm
// end time:	11:43 pm
// ---------------------
// time taken: 	 minutes

type Prac23June01Sol1 struct {
}

func (p Prac23June01Sol1) FindAllOrangesRottenTime(arr [][]int) int {
	rotten := make([]Point, 0)
	var numFresh int

	for r, row := range arr {
		for c := range row {
			switch arr[r][c] {
			case 1:
				numFresh++
			case 2:
				rotten = append(rotten, Point{r, c})
			}
		}
	}

	numRtnOranges, numFreshToRotten := p.calcAllRottenOrangesRottenTime(rotten, arr)
	if numFreshToRotten != numFresh {
		return -1
	}

	return numRtnOranges
}

var June01Move = [4]Point{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func (p Prac23June01Sol1) calcAllRottenOrangesRottenTime(rottenPos []Point, arr [][]int) (int, int) {
	//seen := p.prepareSeenArr(arr)

	q := make([]Point, 0)
	q = append(q, rottenPos...)

	var numFresh int

	var timeTaken int
	for len(q) > 0 {

		//iterate over all the initially present rotten oranges
		for i := len(q); i > 0; i-- {
			rtnPos := q[0]
			q = q[1:]

			for _, m := range June01Move {
				if vPoint, ok := p.isValidMove(rtnPos, m, arr); ok {
					arr[vPoint.r][vPoint.c] = 2 //rotten now
					numFresh++
					q = append(q, *vPoint)
				}
			}
		}

		//only update time new rotten oranges have been added in this iteration
		if len(q) > 0 {
			timeTaken++
		}
	}

	return timeTaken, numFresh
}

func (p Prac23June01Sol1) isValidMove(rtn, move Point, arr [][]int) (*Point, bool) {
	r := rtn.r + move.r
	if r < 0 || r >= len(arr) {
		return nil, false
	}

	c := rtn.c + move.c
	if c < 0 || c >= len(arr[0]) {
		return nil, false
	}

	if arr[r][c] != 1 {
		return nil, false
	}

	return &Point{r, c}, true
}

type Point struct {
	r int
	c int
}
