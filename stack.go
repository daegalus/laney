package laney

// Stack is a LIFO (Last in first out) data structure implementation.
// It is based on a deque container and focuses its API on core
// functionalities: Push, Pop, Head, Size, Empty. Every operations time complexity
// is O(1).
//
// As it is implemented using a Deque container, every operations
// over an instiated Stack are synchronized and safe for concurrent
// usage.
type Stack[T any] struct {
	*Deque[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		Deque: NewDeque[T](),
	}
}

// Push adds on an item on the top of the Stack
func (s *Stack[T]) Push(item T) {
	s.Prepend(item)
}

// Pop removes and returns the item on the top of the Stack
func (s *Stack[T]) Pop() T {
	return s.Shift()
}

// Head returns the item on the top of the stack
func (s *Stack[T]) Head() T {
	return s.First()
}
