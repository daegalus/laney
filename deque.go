package laney

import (
	"container/list"
	"sync"
)

// Deque is a head-tail linked list data structure implementation.
// It is based on a doubly linked list container, so that every
// operations time complexity is O(1).
//
// every operations over an instiated Deque are synchronized and
// safe for concurrent usage.
type Deque[T any] struct {
	sync.RWMutex
	container *list.List
	capacity  int
}

// NewDeque creates a Deque.
func NewDeque[T any]() *Deque[T] {
	return NewCappedDeque[T](-1)
}

// NewCappedDeque creates a Deque with the specified capacity limit.
func NewCappedDeque[T any](capacity int) *Deque[T] {
	return &Deque[T]{
		container: list.New(),
		capacity:  capacity,
	}
}

// Append inserts element at the back of the Deque in a O(1) time complexity,
// returning true if successful or false if the deque is at capacity.
func (s *Deque[T]) Append(item T) bool {
	s.Lock()
	defer s.Unlock()

	if s.capacity < 0 || s.container.Len() < s.capacity {
		s.container.PushBack(item)
		return true
	}

	return false
}

// Prepend inserts element at the Deques front in a O(1) time complexity,
// returning true if successful or false if the deque is at capacity.
func (s *Deque[T]) Prepend(item T) bool {
	s.Lock()
	defer s.Unlock()

	if s.capacity < 0 || s.container.Len() < s.capacity {
		s.container.PushFront(item)
		return true
	}

	return false
}

// Pop removes the last element of the deque in a O(1) time complexity
func (s *Deque[T]) Pop() T {
	s.Lock()
	defer s.Unlock()

	var item T
	var lastContainerItem *list.Element = nil

	lastContainerItem = s.container.Back()
	if lastContainerItem != nil {
		item = s.container.Remove(lastContainerItem).(T)
	}

	return item
}

// Shift removes the first element of the deque in a O(1) time complexity
func (s *Deque[T]) Shift() T {
	s.Lock()
	defer s.Unlock()

	var item T
	var firstContainerItem *list.Element = nil

	firstContainerItem = s.container.Front()
	if firstContainerItem != nil {
		item = s.container.Remove(firstContainerItem).(T)
	}

	return item
}

// First returns the first value stored in the deque in a O(1) time complexity
func (s *Deque[T]) First() T {
	s.RLock()
	defer s.RUnlock()

	item := s.container.Front()
	if item != nil {
		return item.Value.(T)
	} else {
		var nothing T
		return nothing
	}
}

// Last returns the last value stored in the deque in a O(1) time complexity
func (s *Deque[T]) Last() T {
	s.RLock()
	defer s.RUnlock()

	item := s.container.Back()
	if item != nil {
		return item.Value.(T)
	} else {
		var nothing T
		return nothing
	}
}

// Size returns the actual deque size
func (s *Deque[T]) Size() int {
	s.RLock()
	defer s.RUnlock()

	return s.container.Len()
}

// Capacity returns the capacity of the deque, or -1 if unlimited
func (s *Deque[T]) Capacity() int {
	s.RLock()
	defer s.RUnlock()
	return s.capacity
}

// Empty checks if the deque is empty
func (s *Deque[T]) Empty() bool {
	s.RLock()
	defer s.RUnlock()

	return s.container.Len() == 0
}

// Full checks if the deque is full
func (s *Deque[T]) Full() bool {
	s.RLock()
	defer s.RUnlock()

	return s.capacity >= 0 && s.container.Len() >= s.capacity
}
