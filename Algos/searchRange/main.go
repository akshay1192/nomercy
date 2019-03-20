package main

import (
	"fmt"
	"github.com/akshay1192/nomercy/Algos/searchRange/helperFunc"
)

func main() {

	fmt.Println(helperFunc.SearchForRange([]int{0, 1, 21, 33, 45, 45, 45, 45, 61, 71, 73}, 45))
	fmt.Println(helperFunc.SearchForRangeIfElse([]int{0, 1, 21, 33, 45, 45, 45, 45, 61, 71, 73}, 45))

}
