package heap

import (
	"container/heap"
	"errors"
)

// Item is an item that can be added to the priority queue.
type Item interface {
	// Compare returns a bool that can be used to determine
	// ordering in the priority queue.
	Compare(other Item) bool
}

type Heap struct {
	itemHeap *itemHeap
}

func NewHeap() Heap {
	return Heap{
		itemHeap: &itemHeap{},
	}
}

func (p *Heap) Len() int {
	return p.itemHeap.Len()
}

// Insert inserts a new element into the queue. Duplicate Item is allowed
func (p *Heap) Insert(v Item) {
	heap.Push(p.itemHeap, v)
}

// Pop removes the element with the highest priority from the queue and returns it.
// In case of an empty queue, an error is returned.
func (p *Heap) Pop() (Item, error) {
	if len(*p.itemHeap) == 0 {
		return nil, errors.New("empty queue")
	}

	return heap.Pop(p.itemHeap).(Item), nil
}

// Peek returns the element with the highest priority from the queue.
// In case of an empty queue, an error is ret
func (p *Heap) Peek() (Item, error) {
	if len(*p.itemHeap) == 0 {
		return nil, errors.New("empty queue")
	}

	return (*p.itemHeap)[0], nil
}

type itemHeap []Item

func (ih *itemHeap) Len() int {
	return len(*ih)
}

func (ih *itemHeap) Less(i, j int) bool {
	return (*ih)[i].Compare((*ih)[j])
}

func (ih *itemHeap) Swap(i, j int) {
	(*ih)[i], (*ih)[j] = (*ih)[j], (*ih)[i]
}

func (ih *itemHeap) Push(x any) {
	*ih = append(*ih, x.(Item))
}

func (ih *itemHeap) Pop() any {
	old := *ih
	item := old[len(old)-1]
	*ih = old[0 : len(old)-1]
	return item
}
