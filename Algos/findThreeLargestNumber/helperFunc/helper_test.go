package helperFunc

import (
	"reflect"
	"testing"
)

func TestFindThreeLargestNumbers(t *testing.T) {
	actual := FindThreeLargestNumbers([]int{55, 7, 8})
	expected := []int{7, 8, 55}

	if !reflect.DeepEqual(actual, expected) {
		t.Fail()
	}
}

func TestFindThreeLargestNumbersBubbleSort(t *testing.T) {
	actual := FindThreeLargestNumbers([]int{1, 3, 11, 9, 6, 5, 8})
	expected := []int{8, 9, 11}

	if !reflect.DeepEqual(actual, expected) {
		t.Fail()
	}
}
