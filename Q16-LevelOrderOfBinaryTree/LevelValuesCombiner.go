package Q16_LevelOrderOfBinaryTree

import (
	"github.com/solve_common_problems_using_go"
)

type LevelValuesCombiner interface {
	combineValuesAtSameLevel(n *solve_common_problems_using_go.BSTNode) [][]int
}
