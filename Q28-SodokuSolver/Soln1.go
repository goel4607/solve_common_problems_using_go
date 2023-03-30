package Q28_SodokuSolver

type Soln1 struct {
}

func (s *Soln1) init(p [][]int, rowsMap, colsMap, mtxMap *[9]map[int]bool) {
	for i, row := range p {
		if rowsMap[i] == nil {
			rowsMap[i] = make(map[int]bool)
		}

		for j, v := range row {
			if colsMap[j] == nil {
				colsMap[j] = make(map[int]bool)
			}

			if p[i][j] != 0 {
				rowsMap[i][v] = true
				colsMap[j][v] = true

				matIdx := s.mapMatrixId(i, j)
				if mtxMap[matIdx] == nil {
					mtxMap[matIdx] = make(map[int]bool)
				}

				if _, ok := mtxMap[matIdx][v]; !ok {
					mtxMap[matIdx][v] = true
				}
			}
		}
	}
}

func (s *Soln1) SodukuSolver(p [][]int) [][]int {
	var rowsMap, colsMap, mtxMap [9]map[int]bool
	s.init(p, &rowsMap, &colsMap, &mtxMap)

	s.solveSoduku(0, 0, p, &rowsMap, &colsMap, &mtxMap)
	return p
}

func (s *Soln1) solveSoduku(r, c int, p [][]int, rowsMap, colsMap, mtxMap *[9]map[int]bool) bool {
	if r == len(p) || c == len(p[0]) {
		return true
	} else {
		if p[r][c] == 0 {
			for k := 1; k <= 9; k++ {
				p[r][c] = k
				if s.isValid(k, r, c, rowsMap, colsMap, mtxMap) {
					s.update(true, k, r, c, rowsMap, colsMap, mtxMap)

					if c+1 == len(p[0]) { // last column for the row
						if s.solveSoduku(r+1, 0, p, rowsMap, colsMap, mtxMap) {
							return true
						}
					} else {
						if s.solveSoduku(r, c+1, p, rowsMap, colsMap, mtxMap) {
							return true
						}
					}

					s.update(false, k, r, c, rowsMap, colsMap, mtxMap)
				}

				p[r][c] = 0 //revert back
			}
		} else {
			if c+1 == len(p[0]) { // last column for the row
				if s.solveSoduku(r+1, 0, p, rowsMap, colsMap, mtxMap) {
					return true
				}
			} else {
				if s.solveSoduku(r, c+1, p, rowsMap, colsMap, mtxMap) {
					return true
				}
			}
		}
	}

	return false
}

func (s *Soln1) update(mark bool, v int, i, j int, rowsMap, colsMap, mtxMap *[9]map[int]bool) {
	if rowsMap[i] == nil {
		rowsMap[i] = make(map[int]bool)
	}
	if mark {
		rowsMap[i][v] = mark
	} else {
		delete(rowsMap[i], v)
	}

	if colsMap[j] == nil {
		colsMap[j] = make(map[int]bool)
	}
	if mark {
		colsMap[j][v] = mark
	} else {
		delete(colsMap[j], v)
	}

	matIdx := s.mapMatrixId(i, j)
	if mtxMap[matIdx] == nil {
		mtxMap[matIdx] = make(map[int]bool)
	}
	if mark {
		mtxMap[matIdx][v] = mark
	} else {
		delete(mtxMap[matIdx], v)
	}
}

func (s *Soln1) isValid(v int, i, j int, rowsMap, colsMap, mtxMap *[9]map[int]bool) bool {
	if _, ok := rowsMap[i][v]; ok {
		return false
	}

	if _, ok := colsMap[j][v]; ok {
		return false
	}

	matIdx := s.mapMatrixId(i, j)
	if _, ok := mtxMap[matIdx][v]; ok {
		return false
	}

	return true
}

func (s Soln1) mapMatrixId(i, j int) int {
	matRowIdx := i / 3
	matColIdx := j / 3
	return matRowIdx*3 + matColIdx
}
