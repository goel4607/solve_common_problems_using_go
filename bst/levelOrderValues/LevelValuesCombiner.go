package levelOrderValues

import "github.com/solve_common_problems_using_go/bst"

type LevelValuesCombiner interface {
	combineValuesAtSameLevel(n *bst.Node) [][]int
}
