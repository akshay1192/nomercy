// Given an array, find the pair with max difference such that larger number comes after smaller number
// Example 1 : [2 3 10 2 4 7] => 8 (2,10)
// Example 2 : [7 9 6 4 3 1] => 2 (7,9)

package main

import (
	"fmt"
	"math"
)

func main() {

	var size int
	fmt.Scanln(&size)

	arr := []int{}
	for i := 0; i < size; i++ {
		var num int
		fmt.Scanln(&num)
		arr = append(arr, num)
	}

	maxDiff := math.MinInt16
	minEle := arr[0]

	for _, v := range arr {
		if v > minEle && v-minEle > maxDiff {
			maxDiff = v - minEle
		}

		if minEle > v {
			minEle = v
		}
	}

	fmt.Println(maxDiff)
}
