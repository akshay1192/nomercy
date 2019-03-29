// This example demonstrates a Priority queue built using the heap interface.
package util

import (
	"errors"
)

const (
	EmptyHeap = "heap is empty"
)

// An Item is something we manage in a Priority queue.
type Item struct {
	Priority int64 // The Priority of the item in the queue.
	Index    int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, Priority so we use greater than here.
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func GetMaxFromHeap(heap PriorityQueue) (int64,error){
	if len(heap) != 0 {
		return heap[0].Priority,nil
	}

	return -1,errors.New(EmptyHeap)
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in Priority order.
/*func main() {
	// Some items and their priorities.
	nums := []int{8}

	// Create a Priority queue, put the items in it, and
	// establish the Priority queue (heap) invariants.
	pq := make(PriorityQueue, len(nums))
	i := 0
	for _, priority := range nums {
		pq[i] = &Item{
			Priority: priority,
			Index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its Priority.
	item := &Item{
		Priority: 5,
	}
	heap.Push(&pq, item)

	fmt.Println(pq[0].Priority)

	// Insert a new item and then modify its Priority.
	item = &Item{
		Priority: 7,
	}
	heap.Push(&pq, item)

	fmt.Println(pq[0].Priority)
	// Insert a new item and then modify its Priority.
	item = &Item{
		Priority: 10,
	}
	heap.Push(&pq, item)

	fmt.Println(pq[0].Priority)


	// Take the items out; they arrive in decreasing Priority order.
	/*for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.Priority)
	}
}*/