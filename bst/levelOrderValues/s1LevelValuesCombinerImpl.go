package levelOrderValues

import (
	"github.com/solve_common_problems_using_go/bst"
)

type S1LevelValuesCombinerImpl struct {
}

func (s1 S1LevelValuesCombinerImpl) combineValuesAtSameLevel(n *bst.Node) [][]int {
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

func (s1 S1LevelValuesCombinerImpl) combineValuesAtSameLevelInorder(n *bst.Node, lvlValsMap map[int][]int, lvl int) {
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
