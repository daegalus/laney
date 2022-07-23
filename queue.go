package laney

// Queue is a FIFO (First in first out) data structure implementation.
// It is based on a deque container and focuses its API on core
// functionalities: Enqueue, Dequeue, Head, Size, Empty. Every operations time complexity
// is O(1).
//
// As it is implemented using a Deque container, every operations
// over an instiated Queue are synchronized and safe for concurrent
// usage.
type Queue[T any] struct {
	*Deque[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		Deque: NewDeque[T](),
	}
}

// Enqueue adds an item at the back of the queue
func (q *Queue[T]) Enqueue(item T) {
	q.Prepend(item)
}

// Dequeue removes and returns the front queue item
func (q *Queue[T]) Dequeue() T {
	return q.Pop()
}

// Head returns the front queue item
func (q *Queue[T]) Head() T {
	return q.Last()
}
