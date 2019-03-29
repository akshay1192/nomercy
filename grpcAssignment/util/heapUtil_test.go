package util

import (
	"container/heap"
	"testing"
)

// testing with no elements in heap
func TestGetMaxFromHeap_Test1(t *testing.T) {

	var pq PriorityQueue

	actual,err := GetMaxFromHeap(pq)
	expected := -1
	if err.Error() == EmptyHeap && actual == int64(expected){
		actual = -1
	} else {
		t.Fail()
	}

}

// testing heapify function
func TestGetMaxFromHeap_Test2(t *testing.T) {

	var pq PriorityQueue
	pq.Push(&Item{1,0})
	pq.Push(&Item{2,1})

	// heapify
	heap.Init(&pq)

	actual,err := GetMaxFromHeap(pq)
	expected := 2
	if err == nil && actual != int64(expected){
		t.Fail()
	}

}

// testing heapify function and then inserting 11 as incoming num
func TestGetMaxFromHeap_Test3(t *testing.T) {

	var pq PriorityQueue
	pq.Push(&Item{1,0})
	pq.Push(&Item{8,1})
	pq.Push(&Item{3,2})
	pq.Push(&Item{6,3})

	// calling heapify over current elements
	heap.Init(&pq)

	// inserting the number in present heap
	heap.Push(&pq,&Item{10,4})

	actual,err := GetMaxFromHeap(pq)
	expected := 10
	if err == nil && actual != int64(expected){
		t.Fail()
	}

}
