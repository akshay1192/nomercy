package main

import (
	"fmt"
)

type node struct {
	value int
	next  *node
}

func (root *node) addFront(val int) *node {

	newNode := &node{val, nil}
	newNode.next = root
	return newNode

}

func (root *node) addLast(val int) {

	temp := root
	for temp.next != nil {
		temp = temp.next
	}

	newNode := &node{val, nil}
	temp.next = newNode
}

func (root *node) delete(val int) *node {

	temp := root
	var prev *node = nil
	found := false

	for temp != nil {
		if temp.value == val {
			found = true
			break
		} else {
			prev = temp
			temp = temp.next

		}
	}

	if found {
		if temp == root {
			root = root.next
			temp.next = nil
			return root
		} else {
			prev.next = temp.next
		}
	} else {
		fmt.Printf("No node with value %d found \n", val)
	}

	return root
}

func main() {
	root := &node{10, nil}
	root.addLast(20)
	root.addLast(30)
	root.addLast(40)

	root = root.delete(10)
	root = root.addFront(50)
	root = root.delete(50)

	for root != nil {
		fmt.Println(root.value)
		root = root.next
	}
}
