package heap

import (
	"container/heap"
	"errors"
)

// Item is an item that can be added to the priority queue.
type Item interface {
	// Less returns a bool that can be used to determine
	// ordering in the priority queue.
	Less(other Item) bool
}

type Heap[T Item] struct {
	itemHeap *itemHeap[T]
}

func NewHeap[T Item]() Heap[T] {
	return Heap[T]{
		itemHeap: &itemHeap[T]{},
	}
}

func (p *Heap[T]) Len() int {
	return p.itemHeap.Len()
}

// Insert inserts a new element into the queue. Duplicate Item is allowed
func (p *Heap[T]) Insert(v T) {
	heap.Push(p.itemHeap, v)
}

// Pop removes the element with the highest priority from the queue and returns it.
// In case of an empty queue, an error is returned.
func (p *Heap[T]) Pop() (T, error) {
	if len(*p.itemHeap) == 0 {
		var zeroVal T
		return zeroVal, errors.New("empty queue")
	}

	return heap.Pop(p.itemHeap).(T), nil
}

// Peek returns the element with the highest priority from the queue.
// In case of an empty queue, an error is ret
func (p *Heap[T]) Peek() (T, error) {
	if len(*p.itemHeap) == 0 {
		var zeroVal T

		return zeroVal, errors.New("empty queue")
	}

	return (*p.itemHeap)[0], nil
}

type itemHeap[T Item] []T

func (ih *itemHeap[T]) Len() int {
	return len(*ih)
}

func (ih *itemHeap[T]) Less(i, j int) bool {
	return (*ih)[i].Less((*ih)[j])
}

func (ih *itemHeap[T]) Swap(i, j int) {
	(*ih)[i], (*ih)[j] = (*ih)[j], (*ih)[i]
}

func (ih *itemHeap[T]) Push(x any) {
	*ih = append(*ih, x.(T))
}

func (ih *itemHeap[T]) Pop() any {
	old := *ih
	item := old[len(old)-1]
	*ih = old[0 : len(old)-1]
	return item
}
