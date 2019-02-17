package main

import (
	"fmt"
	"github.com/akshay1192/nomercy/Algos/fibonacci/helperFunc"
)

func main() {
	//fmt.Println(helperFunc.CalculateFibonacciHash(7))
	//fmt.Println(helperFunc.CalculateFibonacciHash2(7))
	fmt.Println(helperFunc.CalculateFibonacciWithoutHash(7))

	fmt.Println(helperFunc.CalculateFibonacciLoop(7))
}
