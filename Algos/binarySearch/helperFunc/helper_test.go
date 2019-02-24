package helperFunc

import "testing"

func TestBinarySearchRecursion(t *testing.T) {
	actual := BinarySearchRecursion([]int{1, 2, 3, 4}, 6)
	expected := -1

	if actual != expected {
		t.Fail()
	}
}

func TestBinarySearchRecursion2(t *testing.T) {
	actual := BinarySearchRecursion([]int{-6, -1, 0, 1, 2, 3, 4}, -1)
	expected := 1

	if actual != expected {
		t.Fail()
	}
}

func TestBinarySearchIterative(t *testing.T) {
	actual := BinarySearchIterative([]int{1, 2, 3, 4}, 6)
	expected := -1

	if actual != expected {
		t.Fail()
	}
}

func TestBinarySearchIterative2(t *testing.T) {
	actual := BinarySearchIterative([]int{-6, -1, 0, 1, 2, 3, 4}, -1)
	expected := 1

	if actual != expected {
		t.Fail()
	}
}
