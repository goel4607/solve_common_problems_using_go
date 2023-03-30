package Q12_ImplQueueWithStacks

type Queue interface {
	enqueue(v int)
	dequeue() (int, bool)
	peek() (int, bool)
	empty() bool
}
