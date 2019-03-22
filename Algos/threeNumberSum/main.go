package main

import (
	"fmt"
	"github.com/akshay1192/nomercy/Algos/threeNumberSum/helperFunc"
)

func main() {

	fmt.Println(helperFunc.ThreeNumberSum([]int{12, 3, 1, 2, -6, 5, -8, 6}, 0))
	fmt.Println(helperFunc.ThreeNumberSumNaiveApproach([]int{12, 3, 1, 2, -6, 5, -8, 6}, 0))
}
