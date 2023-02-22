package rightSideView

import (
	"github.com/solve_common_problems_using_go/bst"
)

type RSVSolutionUsingDFS struct {
}

func (s RSVSolutionUsingDFS) combineRightSideView(n *bst.Node) []int {
	if n == nil {
		return []int{}
	}

	valByLvlMap := make(map[int]int)
	s.visitPreorderWithRFirst(n, 0, valByLvlMap)

	values := make([]int, len(valByLvlMap), len(valByLvlMap))
	for i, v := range valByLvlMap {
		values[i] = v
	}

	return values
}

func (s RSVSolutionUsingDFS) visitPreorderWithRFirst(n *bst.Node, level int, valByLvlMap map[int]int) {
	if n == nil {
		return
	}

	if _, ok := valByLvlMap[level]; !ok {
		valByLvlMap[level] = n.Data
	}

	if n.Right != nil {
		s.visitPreorderWithRFirst(n.Right, level+1, valByLvlMap)
	}

	if n.Left != nil {
		s.visitPreorderWithRFirst(n.Left, level+1, valByLvlMap)
	}

	return
}
