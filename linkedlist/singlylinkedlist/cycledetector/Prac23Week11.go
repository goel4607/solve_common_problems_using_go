package cycledetector

type Prac23Week11 struct {
}

func (p Prac23Week11) detectCycle(head *CNode) bool {
	slow := head
	fast := head.Next

	for slow != fast {
		slow = slow.Next
		if fast == nil || fast.Next == nil {
			return false
		} else {
			fast = fast.Next.Next
		}
	}

	return true
}
