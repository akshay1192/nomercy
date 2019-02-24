package helperFunc

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	actual := BubbleSort([]int{3, 2, 1, -1, 0, 8})
	expected := []int{-1, 0, 1, 2, 3, 8}

	if !reflect.DeepEqual(actual, expected) {
		t.Fail()
	}
}

func TestBubbleSort2(t *testing.T) {
	actual := BubbleSort([]int{5, 4, 3, 2, 1})
	expected := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(actual, expected) {
		t.Fail()
	}
}

func BenchmarkBubbleSort(b *testing.B) {

	for i := 0; i < b.N; i++ {
		BubbleSort([]int{3, 2, 1, -1, 0, 8})
	}
}
