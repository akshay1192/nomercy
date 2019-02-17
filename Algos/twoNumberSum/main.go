package main

import (
	"fmt"
	"github.com/akshay1192/nomercy/Algos/twoNumberSum/helperFunc"
	"time"
)

func main() {

	startTime := time.Now()
	fmt.Println("Result : ", helperFunc.CheckSum([]int{8, 10, 12, 13, 14, 2, 3, 4, 5}, 9))
	elapsedTime := time.Since(startTime)

	fmt.Println("Total time : ", elapsedTime)
}
