package bst

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
