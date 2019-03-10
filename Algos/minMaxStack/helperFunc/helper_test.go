package helperFunc

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewMinMaxStack()

	stack.Push(2)
	if stack.GetMin() != 2 || stack.GetMax() != 2 || stack.Peek() != 2 {
		t.Fail()
	}

	stack.Push(7)
	if stack.GetMin() != 2 || stack.GetMax() != 7 || stack.Peek() != 7 {
		t.Fail()
	}

	stack.Push(1)
	if stack.GetMin() != 1 || stack.GetMax() != 7 || stack.Peek() != 1 {
		t.Fail()
	}

	stack.Push(8)
	if stack.GetMin() != 1 || stack.GetMax() != 8 || stack.Peek() != 8 {
		t.Fail()
	}

	stack.Push(3)
	if stack.GetMin() != 1 || stack.GetMax() != 8 || stack.Peek() != 3 {
		t.Fail()
	}

	stack.Push(9)
	if stack.GetMin() != 1 || stack.GetMax() != 9 || stack.Peek() != 9 {
		t.Fail()
	}

	stack.Pop()
	if stack.GetMin() != 1 || stack.GetMax() != 8 || stack.Peek() != 3 {
		t.Fail()
	}
}
