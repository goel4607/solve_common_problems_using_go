package Q16_LevelOrderOfBinaryTree

import (
	"github.com/solve_common_problems_using_go"
)

type SolutionUsingDFS struct {
}

func (s1 SolutionUsingDFS) combineValuesAtSameLevel(n *solve_common_problems_using_go.BSTNode) [][]int {
	if n == nil {
		return nil
	}

	// each entry represents a level so at each level, the array contains all the elements present at that level
	lvlValsMap := make(map[int][]int)

	s1.combineValuesAtSameLevelInorder(n, lvlValsMap, 1)

	lvlValsArr := make([][]int, len(lvlValsMap))
	for lvl, vals := range lvlValsMap {
		lvlValsArr[lvl-1] = vals
	}

	return lvlValsArr
}

func (s1 SolutionUsingDFS) combineValuesAtSameLevelInorder(n *solve_common_problems_using_go.BSTNode, lvlValsMap map[int][]int, lvl int) {
	if n != nil {
		if lvlVals, ok := lvlValsMap[lvl]; !ok {
			lvlVals := make([]int, 0)
			lvlVals = append(lvlVals, n.Data)

			lvlValsMap[lvl] = lvlVals
		} else {
			lvlVals = append(lvlVals, n.Data)
			lvlValsMap[lvl] = lvlVals
		}
	}

	if n.Left != nil {
		s1.combineValuesAtSameLevelInorder(n.Left, lvlValsMap, lvl+1)
	}

	if n.Right != nil {
		s1.combineValuesAtSameLevelInorder(n.Right, lvlValsMap, lvl+1)
	}
}
