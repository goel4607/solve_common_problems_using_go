package levelOrderValues

import (
	"github.com/solve_common_problems_using_go/bst"
)

type SolutionUsingBFS struct {
}

func (s2 SolutionUsingBFS) combineValuesAtSameLevel(n *bst.Node) [][]int {
	lvlValuesMap := make(map[int][]int, 0)

	q := Queue{}
	q.enqueue(n)

	lvl := 0
	for !q.empty() {

		var count int
		maxLen := q.len()

		values := make([]int, 0, maxLen)
		for count < maxLen {
			t := q.dequeue()

			values = append(values, t.Data)
			if t.Left != nil {
				q.enqueue(t.Left)
			}

			if t.Right != nil {
				q.enqueue(t.Right)
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

type Queue struct {
	vals []*bst.Node
}

func (q *Queue) enqueue(n *bst.Node) {
	q.vals = append(q.vals, n)
}

func (q *Queue) dequeue() *bst.Node {
	v := q.vals[0]
	q.vals = q.vals[1:]
	return v
}

func (q *Queue) empty() bool {
	if len(q.vals) == 0 {
		return true
	} else {
		return false
	}
}

func (q *Queue) peek() *bst.Node {
	if q.empty() {
		return nil
	}

	return q.vals[0]
}

func (q *Queue) len() int {
	return len(q.vals)
}
