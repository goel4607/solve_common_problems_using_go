package useStackAsQueue

type Solution1 struct {
	values []int
}

func (s1 *Solution1) enqueue(v int) {
	s1.values = append(s1.values, v)
}

func (s1 *Solution1) dequeue() (int, bool) {
	if s1.empty() {
		return 0, false
	}

	v := s1.values[0]         // keep the first element
	s1.values = s1.values[1:] // remove the element at first position
	return v, true
}

func (s1 *Solution1) empty() bool {
	if len(s1.values) == 0 {
		return true
	} else {
		return false
	}
}

func (s1 *Solution1) peek() (int, bool) {
	if s1.empty() {
		return 0, false
	}

	return s1.values[0], true
}
