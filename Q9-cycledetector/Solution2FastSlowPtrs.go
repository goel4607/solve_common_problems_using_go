package Q9_cycledetector

type Solution2FastSlowPtrs struct {
}

func (s2 Solution2FastSlowPtrs) detectCycle(h *CNode) bool {
	if h == nil || h.Next == nil {
		return false
	}

	slow := h
	fast := h.Next

	for slow != fast {
		slow = slow.Next
		if fast != nil && fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}
	}

	return true
}
