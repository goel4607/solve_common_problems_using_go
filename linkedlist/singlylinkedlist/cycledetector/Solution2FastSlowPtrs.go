package cycledetector

type Solution2FastSlowPtrs struct {
}

func (s2 Solution2FastSlowPtrs) detectCycle(h *CNode) bool {
	if h == nil {
		return false
	}

	slow := h
	fast := h.Next

	for slow != fast {
		slow = slow.Next
		if fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}
	}

	return true
}
