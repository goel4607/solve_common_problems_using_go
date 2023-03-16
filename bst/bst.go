package bst

import "math/rand"

type Node struct {
	Left  *Node
	Data  int
	Right *Node
}

func (n *Node) Insert(v int) {
	if n == nil {
		return
	}

	if v < n.Data {
		if n.Left == nil {
			n.Left = &Node{Data: v}
		} else {
			n.Left.Insert(v)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Data: v}
		} else {
			n.Right.Insert(v)
		}
	}
}

func (n *Node) InsertInRandomOrder(v int) {
	if n == nil {
		return
	}

	r := rand.Intn(2) // 0 for left and 1 for right
	if r == 0 && v < n.Data {
		if n.Left == nil {
			n.Left = &Node{Data: v}
		} else {
			n.Left.InsertInRandomOrder(v)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Data: v}
		} else {
			n.Right.InsertInRandomOrder(v)
		}
	}
}

func NewWithValuesInSpecifiedOrder(vals []int) (n *Node) {
	if len(vals) == 0 {
		return
	}

	v := vals[0]    // get the top value from the Q
	vals = vals[1:] // remove the top value from the Q

	head := &Node{Data: v}

	q := Queue{}
	q.Enqueue(head)

	for len(vals) > 0 {
		v := vals[0]    // get the top value from the Q
		vals = vals[1:] // remove the top value from the Q
		n := q.Dequeue()
		n.Left = &Node{Data: v}
		q.Enqueue(n.Left)

		if len(vals) > 0 {
			v = vals[0]     // get the top value from the Q
			vals = vals[1:] // remove the top value from the Q
			n.Right = &Node{Data: v}
			q.Enqueue(n.Right)
		}
	}

	return head
}

func (n *Node) Visit() int {
	if n != nil {
		return n.Data
	}

	return -1
}

func (n *Node) SortedList() []int {
	all := make([]int, 0)
	return n.inorder(all)
}

func (n *Node) inorder(all []int) []int {
	if n == nil {
		return all
	}

	if n.Left != nil {
		all = n.Left.inorder(all)
	}

	all = append(all, n.Data)

	if n.Right != nil {
		all = n.Right.inorder(all)
	}

	return all
}
