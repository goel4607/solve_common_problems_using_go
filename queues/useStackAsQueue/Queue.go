package useStackAsQueue

type Queue interface {
	enqueue(v int)
	dequeue() (int, bool)
	peek() (int, bool)
	empty() bool
}
