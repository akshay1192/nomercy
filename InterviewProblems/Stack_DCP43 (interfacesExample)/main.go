// Implement stack for all data types

package main

import (
	"fmt"
)

type StackIf interface {
	push(element interface{})
	pop() interface{}
	max() interface{}
}

type Stack struct {
	stack    []interface{}
	maxStack []interface{}
}

func (ss *Stack) push(element interface{}) {

	ss.stack = append(ss.stack, element)
	fmt.Println("Pushed : ",element)
	if len(ss.maxStack) == 0 {
		ss.maxStack = append(ss.maxStack, element)
		return
	}

	switch element.(type) {
	case int:
		if element.(int) > ss.maxStack[len(ss.maxStack)-1].(int) {
			ss.maxStack = append(ss.maxStack, element)
		} else {
			ss.maxStack = append(ss.maxStack, ss.maxStack[len(ss.maxStack)-1])
		}
	case string:
		if element.(string) > ss.maxStack[len(ss.maxStack)-1].(string) {
			ss.maxStack = append(ss.maxStack, element)
		} else {
			ss.maxStack = append(ss.maxStack, ss.maxStack[len(ss.maxStack)-1])
		}
	}
}

func (ss *Stack) pop() interface{} {
	if len(ss.stack) == 0 {
		return -1
	}

	poppedElement := ss.stack[len(ss.stack)-1]
	ss.stack = ss.stack[:len(ss.stack)-1]
	ss.maxStack = ss.maxStack[:len(ss.maxStack)-1]

	return poppedElement
}

func (ss *Stack) max() interface{}{
	if len(ss.stack) == 0 {
		return -1
	}
	return ss.maxStack[len(ss.maxStack)-1]
}

func main() {
	intStackObj := Stack{stack : make([]interface{},0),maxStack:make([]interface{},0)}

	intStackObj.push(50)
	intStackObj.push(20)
	intStackObj.pop()
	intStackObj.push(45)
	intStackObj.max()
	intStackObj.push(55)
	intStackObj.max()
	intStackObj.pop()
	intStackObj.max()

	stringStackObj := Stack{stack:make([]interface{},0),maxStack:make([]interface{},0)}
	stringStackObj.push("Akshay")
	stringStackObj.push("Akash")
	stringStackObj.pop()
	stringStackObj.max()
	stringStackObj.push("Gupta")
	stringStackObj.max()

}
