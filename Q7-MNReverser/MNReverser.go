package Q7_MNReverser

import (
	"github.com/solve_common_problems_using_go"
)

type MNReverser interface {
	reverse(head *solve_common_problems_using_go.SLLNode, m, n int) *solve_common_problems_using_go.SLLNode
}
