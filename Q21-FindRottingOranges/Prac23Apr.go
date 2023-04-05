package Q21_FindRottingOranges

const (
	Up1 = iota
	Down1
	Left1
	Right1
)

type Pos1 struct {
	r int
	c int
}

var Moves1 = []Pos1{
	Up1:    {-1, 0},
	Down1:  {1, 0},
	Left1:  {0, -1},
	Right1: {0, 1},
}

const (
	NoOrange1 = iota
	Orange1
	RottenOrange1
)

type Prac23Apr struct {
}

func (s Prac23Apr) findAllOrangesRottenTime(arr [][]int) int {
	if len(arr) == 0 || len(arr[0]) == 0 {
		return 0
	}

	var ttlTime int

	//for i, r := range arr {
	//	for j := range r {
	//		if arr[i][j] == RottenOrange1 {
	//			ttlTime += s.visitBFS(arr, Pos1{i, j})
	//		}
	//	}
	//}

	rPos := s.findAllRottenOrangePositions(arr)
	ttlTime = s.visitBFS(arr, rPos)
	for i, r := range arr {
		for j := range r {
			if arr[i][j] == Orange1 {
				return -1
			}
		}
	}

	return ttlTime
}

func (s Prac23Apr) findAllRottenOrangePositions(arr [][]int) []Pos1 {
	var rPos []Pos1
	for i, r := range arr {
		for j := range r {
			switch arr[i][j] {
			case NoOrange1, Orange1:
				continue
			case RottenOrange1:
				rPos = append(rPos, Pos1{r: i, c: j})
			}
		}
	}

	return rPos
}

func (s Prac23Apr) visitBFS(arr [][]int, rPos []Pos1) int {
	var q []Pos1
	q = append(q, rPos...)

	var ttlTime int

	currQSize := len(q)

	for len(q) > 0 {
		if currQSize == 0 {
			ttlTime++

			currQSize = len(q)
		}
		v := q[0]
		q = q[1:]

		arr[v.r][v.c] = RottenOrange1

		currQSize--

		for _, m := range Moves1 {
			if s.isValidMove1(arr, v, m) {
				q = append(q, Pos1{v.r + m.r, v.c + m.c})
			}
		}
	}

	return ttlTime
}

func (s Prac23Apr) isValidMove1(arr [][]int, p Pos1, m Pos1) bool {
	r := p.r + m.r
	if r < 0 || r >= len(arr) {
		return false
	}

	c := p.c + m.c
	if c < 0 || c >= len(arr[0]) {
		return false
	}

	if arr[r][c] == NoOrange1 {
		return false
	}

	if arr[r][c] == RottenOrange1 {
		return false
	}

	return true
}
