package interfaceDesignMonarchy

import "fmt"

type Node struct {
	name     string
	isDead   bool
	children []*Node
}

func (n *Node) findNode(name string) *Node {
	if n.name == name {
		return n
	}

	for _, c := range n.children {
		found := c.findNode(name)
		if found != nil {
			return found
		}
	}

	return nil
}

func (n *Node) traversePreorder(names []string) []string {
	if !n.isDead {
		names = append(names, n.name)
	}

	for _, c := range n.children {
		names = c.traversePreorder(names)
	}

	return names
}

type Soln1 struct {
	hier Node
}

func (s *Soln1) Birth(child, parent string) error {
	if len(parent) == 0 {
		if len(s.hier.name) == 0 {
			s.hier = Node{name: child}
			return nil
		} else {
			return fmt.Errorf("please specify parent as top monarch is already present!! and his name is=%s", s.hier.name)
		}
	}

	pNode := s.hier.findNode(parent)
	if pNode == nil {
		return fmt.Errorf("parent=%s is not found", parent)
	}

	pNode.children = append(pNode.children, &Node{name: child})
	return nil
}

func (s *Soln1) Death(name string) error {
	n := s.hier.findNode(name)
	if n == nil {
		return fmt.Errorf("name=%s is not found", name)
	}

	n.isDead = true
	return nil
}

func (s *Soln1) GetOrderOfSuccession() []string {
	return s.hier.traversePreorder(make([]string, 0))
}
