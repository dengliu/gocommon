package heap

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type myfloat float64

func (f myfloat) Compare(other Item) bool {
	return f < other.(myfloat)
}

func TestHeap(t *testing.T) {
	pq := NewHeap()

	elements := []myfloat{5, 3, 7, 8, 6, 2, 9}
	for _, e := range elements {
		pq.Insert(e)
	}

	sort.Slice(elements, func(i, j int) bool {
		return elements[i] < elements[j]
	})

	for _, e := range elements {
		item, err := pq.Peek()
		assert.NoError(t, err)

		assert.Equal(t, e, item)

		item, err = pq.Pop()
		assert.NoError(t, err)
		assert.Equal(t, e, item)
	}
}

type myMaxHeapFloat float64

func (f myMaxHeapFloat) Compare(other Item) bool {
	return f > other.(myMaxHeapFloat)
}

func TestMaxHeap(t *testing.T) {
	pq := NewHeap()

	elements := []myMaxHeapFloat{5, 3, 7, 8, 6, 2, 9}
	for _, e := range elements {
		pq.Insert(e)
	}

	sort.Slice(elements, func(i, j int) bool {
		return elements[i] > elements[j]
	})

	for _, e := range elements {
		item, err := pq.Peek()
		assert.NoError(t, err)
		assert.Equal(t, e, item)

		item, err = pq.Pop()
		assert.NoError(t, err)
		assert.Equal(t, e, item)
	}
}

type mydata struct {
	name string
	rank int
}

func (m mydata) Compare(o Item) bool {
	return m.rank < o.(mydata).rank
}

func TestHeapLen(t *testing.T) {
	pq := NewHeap()
	assert.Zero(t, pq.Len(), "empty queue should have length 0")

	mydataSlice := []mydata{{"foo", 1}, {"bar", 3}, {"foobar", 2}}
	for _, d := range mydataSlice {
		pq.Insert(d)
	}

	assert.Equal(t, len(mydataSlice), pq.Len())

	sort.Slice(mydataSlice, func(i, j int) bool {
		return mydataSlice[i].rank < mydataSlice[j].rank
	})

	for _, d := range mydataSlice {
		i, err := pq.Pop()
		assert.NoError(t, err)
		assert.Equal(t, d, i)
	}
}
