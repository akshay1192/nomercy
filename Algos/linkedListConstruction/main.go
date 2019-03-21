package main

import "fmt"

type Node struct {
	value      int
	prev, next *Node
}

type DoublyLinkedList struct {
	Head, Tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (ll *DoublyLinkedList) setHead(node *Node) {

	if ll.Head == nil && ll.Tail == nil {
		ll.Tail = node
		ll.Head = node
		return
	}

	node.next = ll.Head
	ll.Head.prev = node
	ll.Head = node

}

func (ll *DoublyLinkedList) setTail(node *Node) {

	if ll.Tail == nil && ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return
	}

	node.prev = ll.Tail
	ll.Tail.next = node
	ll.Tail = node

}

func (ll *DoublyLinkedList) insertBefore(node, nodeToInsert *Node) {

	fmt.Println("Before : ", node)
	fmt.Println("Node : ", nodeToInsert)

	tail := ll.Tail

	for tail != node {
		tail = tail.prev
	}

	prev := tail.prev
	prev.next = node
	tail.prev = node

	node.prev = prev
	node.next = tail

}

/*func (ll *DoublyLinkedList) insertAfter(node, nodeToInsert *Node) {
	// Write your code here.
}

func (ll *DoublyLinkedList) insertAtPosition(position int, nodeToInsert *Node) {
	// Write your code here.
}

func (ll *DoublyLinkedList) removeNodesWithValue(value int) {
	// Write your code here.
}

func (ll *DoublyLinkedList) remove(node *Node) {
	// Write your code here.
}

func (ll *DoublyLinkedList) containsNodeWithValue(value int) bool {
	// Write your code here.
}*/

func main() {

	dll := DoublyLinkedList{}

	node := Node{10, nil, nil}
	dll.setHead(&node)

	node2 := Node{11, nil, nil}
	dll.setHead(&node2)

	fmt.Println("main : ", &node)

	node3 := Node{15, nil, nil}
	dll.insertBefore(&node, &node3)

	for dll.Head != nil {
		fmt.Println(dll.Head.value)
		dll.Head = dll.Head.next
	}

}
