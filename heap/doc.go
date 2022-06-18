// Package heap provoide a easy to use go heap container pkg by
// implementing the "container/heap" interface.

// Example of using github.com/dengliu/gocommon/heap
/*

import (
	"fmt"

	"github.com/dengliu/gocommon/heap"
)

func main() {
	fmt.Println("hello")

	hp := heap.NewHeap[myItem]()
	hp.Insert(myItem{1.2})
	hp.Insert(myItem{2.3})
	hp.Insert(myItem{3.3})
	hp.Insert(myItem{4.3})
	hp.Insert(myItem{4.3})
	hp.Insert(myItem{3.3})
	hp.Insert(myItem{4.3})

	for hp.Len() > 0 {
		fmt.Println(hp.Peek())
		i, _ := hp.Pop()
		fmt.Println(i)
	}
}

type myItem struct {
	rank float64
	data string
}

func (h myItem) Less(other heap.Item) bool {
	return h.data < other.(myItem).data
}

*/
package heap
