package reverse

import "github.com/solve_common_problems_using_go/linkedlist/singlylinkedlist"

type Reverser interface {
	reverse(n *singlylinkedlist.Node) *singlylinkedlist.Node
}
