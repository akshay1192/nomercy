package helperFunc

import (
	"reflect"
	"testing"
)

func TestCheckSumLoopCase1(t *testing.T) {
	actual := CheckSumLoop([]int{8, 10, 5}, 14)
	expected := []int{}

	if !reflect.DeepEqual(actual, expected) {
		t.Fail()
	}
}

func TestCheckSumLoopCase2(t *testing.T) {
	actual := CheckSumLoop([]int{8, 10, 12, 13, 14, 2, 3, 4, 5}, 9)
	expected := []int{4, 5}

	if !reflect.DeepEqual(actual, expected) {
		t.Error("Failed")
	}
}

func TestCheckSumSortCase1(t *testing.T) {
	actual := CheckSumSort([]int{8, 10, 5}, 14)
	expected := []int{}

	if !reflect.DeepEqual(actual, expected) {
		t.Fail()
	}
}

func TestCheckSumSortCase2(t *testing.T) {
	actual := CheckSumSort([]int{8, 10, 12, 13, 14, 2, 3, 4, 5}, 9)
	expected := []int{4, 5}

	if !reflect.DeepEqual(actual, expected) {
		t.Error("Failed")
	}
}

func TestTargetSumHashCase1(t *testing.T) {

	actual := CheckSumHash([]int{8, 10, 12, 13, 14, 2, 3, 4, 5}, 9)
	expected := []int{4, 5}

	if !reflect.DeepEqual(actual, expected) {
		t.Error("Failed")
	}

}

func TestTargetSumHashCase2(t *testing.T) {

	actual := CheckSumHash([]int{8, 10, 5}, 14)
	expected := []int{}

	if !reflect.DeepEqual(actual, expected) {
		t.Error("Failed")
	}

}

func BenchmarkTargetSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheckSumHash([]int{8, 10, 12, 13, 14, 2, 3, 4, 5}, 9)
	}
}

func benchmarkTargetSumWithCustomInput(inputArr []int, targetSum int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheckSumHash(inputArr, targetSum)
	}
}

func BenchmarkTargetSumWithInput(b *testing.B) {
	benchmarkTargetSumWithCustomInput([]int{1, 2, 3, 4, 5}, 6, b)
}
