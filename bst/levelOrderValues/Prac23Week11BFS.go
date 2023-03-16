package levelOrderValues

import "github.com/solve_common_problems_using_go/bst"

// start time: 9:31 pm
// end time:

type Prac23Week11BFS struct {
}

func (p Prac23Week11BFS) combineValuesAtSameLevel(n *bst.Node) [][]int {
	if n == nil {
		return [][]int{}
	}

	levelValues := make([][]int, 0)

	q := []*bst.Node{}
	q = append(q, n)

	for len(q) > 0 {
		count := len(q)

		vals := make([]int, 0)
		for count > 0 {
			v := q[0]
			q = q[1:]

			vals = append(vals, v.Data)

			if v.Left != nil {
				q = append(q, v.Left)
			}

			if v.Right != nil {
				q = append(q, v.Right)
			}

			count--
		}

		levelValues = append(levelValues, vals)
	}

	return levelValues
}
