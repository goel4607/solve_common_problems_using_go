package Q12_ImplQueueWithStacks

// start time: 12:35 am
// end time: 01:03 am
type Prac23Week11 struct {
	s1 MyStack
	s2 MyStack
}

func (q *Prac23Week11) enqueue(v int) {
	q.s1.push(v)
}

func (q *Prac23Week11) dequeue() (int, bool) {
	if q.s2.isEmpty() {
		if q.s1.isEmpty() {
			return -1, false
		}

		for !q.s1.isEmpty() {
			q.s2.push(q.s1.pop())
		}
	}

	return q.s2.pop(), true
}

func (q *Prac23Week11) peek() (int, bool) {
	if q.s2.isEmpty() {
		if q.s1.isEmpty() {
			return -1, false
		}

		for !q.s1.isEmpty() {
			q.s2.push(q.s1.pop())
		}
	}

	return q.s2.peek(), true
}

func (q *Prac23Week11) empty() bool {
	if q.s1.isEmpty() && q.s2.isEmpty() {
		return true
	}

	return false
}

type MyStack []int

func (s *MyStack) push(v int) {
	*s = append(*s, v)
}

func (s *MyStack) pop() int {
	origLen := len(*s)
	v := (*s)[origLen-1]
	*s = (*s)[:origLen-1]
	return v
}

func (s *MyStack) isEmpty() bool {
	if len(*s) == 0 {
		return true
	}

	return false
}

func (s *MyStack) peek() int {
	return (*s)[len(*s)-1]
}
