package Q7_LinkedListReverse

import (
	"github.com/solve_common_problems_using_go"
)

type Reverser interface {
	reverse(n *solve_common_problems_using_go.SLLNode) *solve_common_problems_using_go.SLLNode
}
