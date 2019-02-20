package main

import (
	"fmt"
	"github.com/akshay1192/nomercy/Algos/fibonacci/helperFunc"
	"time"
)

func main() {

	startTime := time.Now()

	fmt.Println(helperFunc.CalculateFibonacciWithoutHash(7))
	fmt.Println(helperFunc.CalculateFibonacciHash(7))
	fmt.Println(helperFunc.CalculateFibonacciLoop(7))

	elapsedTime := time.Since(startTime)

	fmt.Printf("Time elapsed : %v\n", elapsedTime)

}
