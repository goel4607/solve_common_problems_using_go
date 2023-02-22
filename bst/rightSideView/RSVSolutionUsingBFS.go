package rightSideView

import "github.com/solve_common_problems_using_go/bst"

type RSVSolutionUsingBFS struct {
}

func (s1 RSVSolutionUsingBFS) combineRightSideView(n *bst.Node) []int {

	q := bst.Queue{}
	q.Enqueue(n)

	var values []int

	for !q.Empty() {

		var count int
		maxCount := q.Len()

		var lastVal int // last value is always the right hand node of the tree
		for count < maxCount {
			t := q.Dequeue()
			lastVal = t.Data

			if t.Left != nil {
				q.Enqueue(t.Left)
			}

			if t.Right != nil {
				q.Enqueue(t.Right)
			}

			count++
		}

		values = append(values, lastVal)
	}

	return values
}
