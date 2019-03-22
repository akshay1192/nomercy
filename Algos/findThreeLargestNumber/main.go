package main

import (
	"fmt"
	"github.com/akshay1192/nomercy/Algos/findThreeLargestNumber/helperFunc"
)

func main() {
	fmt.Println(helperFunc.FindThreeLargestNumbersBubbleSort([]int{1, 3, 11, 9, 6, 5, 8}))
	fmt.Println(helperFunc.FindThreeLargestNumbers([]int{1, 3, 11, 9, 6, 5, 8}))

}
