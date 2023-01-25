package mnReverser

import "github.com/solve_common_problems_using_go/linkedlist/singlylinkedlist"

type MNReverser interface {
	reverse(head *singlylinkedlist.Node, m, n int) *singlylinkedlist.Node
}
