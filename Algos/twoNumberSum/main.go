package main

import (
	"fmt"
	"github.com/akshay1192/nomercy/Algos/twoNumberSum/helperFunc"
	"time"
)

func main() {

	startTime := time.Now()
	fmt.Println("Result from hash : ", helperFunc.CheckSumHash([]int{8, 10, 12, 13, 14, 2, 3, 4, 5}, 9))
	fmt.Println("Result from sort : ", helperFunc.CheckSumSort([]int{8, 10, 12, 13, 14, 2, 3, 4, 5}, 9))
	fmt.Println("Result from loop : ", helperFunc.CheckSumLoop([]int{8, 10, 12, 13, 14, 2, 3, 4, 5}, 9))
	elapsedTime := time.Since(startTime)

	fmt.Println("Total time : ", elapsedTime)
}
