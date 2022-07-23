package laney

import (
	"fmt"
	"sync"
)

// PQType represents a priority queue ordering kind (see MAXPQ and MINPQ)
type PQType int

const (
	MAXPQ PQType = iota
	MINPQ
)

type item[T any] struct {
	value    T
	priority int
}

// PQueue is a heap priority queue data structure implementation.
// It can be whether max or min ordered and it is synchronized
// and is safe for concurrent operations.
type PQueue[T any] struct {
	sync.RWMutex
	items      []*item[T]
	elemsCount int
	comparator func(int, int) bool
}

func newItem[T any](value T, priority int) *item[T] {
	return &item[T]{
		value:    value,
		priority: priority,
	}
}

func (i *item[T]) String() string {
	return fmt.Sprintf("<item value:%v priority:%d>", i.value, i.priority)
}

// NewPQueue creates a new priority queue with the provided pqtype
// ordering type
func NewPQueue[T any](pqType PQType) *PQueue[T] {
	var cmp func(int, int) bool

	if pqType == MAXPQ {
		cmp = max
	} else {
		cmp = min
	}

	items := make([]*item[T], 1)
	items[0] = nil // Heap queue first element should always be nil

	return &PQueue[T]{
		items:      items,
		elemsCount: 0,
		comparator: cmp,
	}
}

// Push the value item into the priority queue with provided priority.
func (pq *PQueue[T]) Push(value T, priority int) {
	item := newItem(value, priority)

	pq.Lock()
	pq.items = append(pq.items, item)
	pq.elemsCount += 1
	pq.swim(pq.size())
	pq.Unlock()
}

// Pop and returns the highest/lowest priority item (depending on whether
// you're using a MINPQ or MAXPQ) from the priority queue
func (pq *PQueue[T]) Pop() (T, int) {
	pq.Lock()
	defer pq.Unlock()

	if pq.size() < 1 {
		var nothing T
		return nothing, 0
	}

	var max *item[T] = pq.items[1]

	pq.exch(1, pq.size())
	pq.items = pq.items[0:pq.size()]
	pq.elemsCount -= 1
	pq.sink(1)

	return max.value, max.priority
}

// Head returns the highest/lowest priority item (depending on whether
// you're using a MINPQ or MAXPQ) from the priority queue
func (pq *PQueue[T]) Head() (T, int) {
	pq.RLock()
	defer pq.RUnlock()

	if pq.size() < 1 {
		var nothing T
		return nothing, 0
	}

	headValue := pq.items[1].value
	headPriority := pq.items[1].priority

	return headValue, headPriority
}

// Size returns the elements present in the priority queue count
func (pq *PQueue[T]) Size() int {
	pq.RLock()
	defer pq.RUnlock()
	return pq.size()
}

// Check queue is empty
func (pq *PQueue[T]) Empty() bool {
	pq.RLock()
	defer pq.RUnlock()
	return pq.size() == 0
}

func (pq *PQueue[T]) size() int {
	return pq.elemsCount
}

func max(i, j int) bool {
	return i < j
}

func min(i, j int) bool {
	return i > j
}

func (pq *PQueue[T]) less(i, j int) bool {
	return pq.comparator(pq.items[i].priority, pq.items[j].priority)
}

func (pq *PQueue[T]) exch(i, j int) {
	var tmpItem *item[T] = pq.items[i]

	pq.items[i] = pq.items[j]
	pq.items[j] = tmpItem
}

func (pq *PQueue[T]) swim(k int) {
	for k > 1 && pq.less(k/2, k) {
		pq.exch(k/2, k)
		k = k / 2
	}

}

func (pq *PQueue[T]) sink(k int) {
	for 2*k <= pq.size() {
		var j int = 2 * k

		if j < pq.size() && pq.less(j, j+1) {
			j++
		}

		if !pq.less(k, j) {
			break
		}

		pq.exch(k, j)
		k = j
	}
}
