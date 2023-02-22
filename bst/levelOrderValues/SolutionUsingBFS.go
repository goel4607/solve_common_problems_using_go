package levelOrderValues

import (
	"github.com/solve_common_problems_using_go/bst"
)

type SolutionUsingBFS struct {
}

func (s2 SolutionUsingBFS) combineValuesAtSameLevel(n *bst.Node) [][]int {
	lvlValuesMap := make(map[int][]int, 0)

	q := bst.Queue{}
	q.Enqueue(n)

	lvl := 0
	for !q.Empty() {

		var count int
		maxLen := q.Len()

		values := make([]int, 0, maxLen)
		for count < maxLen {
			t := q.Dequeue()

			values = append(values, t.Data)
			if t.Left != nil {
				q.Enqueue(t.Left)
			}

			if t.Right != nil {
				q.Enqueue(t.Right)
			}

			count++
		}

		lvlValuesMap[lvl] = values
		lvl++
	}

	allLvlValues := make([][]int, len(lvlValuesMap), len(lvlValuesMap))
	for k, v := range lvlValuesMap {
		allLvlValues[k] = v
	}

	return allLvlValues
}
