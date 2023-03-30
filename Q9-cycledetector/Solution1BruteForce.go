package Q9_cycledetector

type Solution1BruteForce struct {
}

func (s1 Solution1BruteForce) detectCycle(head *CNode) bool {
	seenData := make(map[any]bool)

	h := head
	for h != nil {
		if _, ok := seenData[h.Data]; ok {
			return true
		}

		seenData[h.Data] = true
		h = h.Next
	}

	return false
}
