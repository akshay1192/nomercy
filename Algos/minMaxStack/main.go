package main

import (
	"fmt"
	"github.com/akshay1192/nomercy/Algos/minMaxStack/helperFunc"
)

func main() {
	stack := helperFunc.NewMinMaxStack()

	stack.Push(2)
	fmt.Println(stack.GetMin(),stack.GetMax(),stack.Peek())
	stack.Push(7)
	fmt.Println(stack.GetMin(),stack.GetMax(),stack.Peek())
	stack.Push(1)
	fmt.Println(stack.GetMin(),stack.GetMax(),stack.Peek())
	stack.Push(8)
	fmt.Println(stack.GetMin(),stack.GetMax(),stack.Peek())
	stack.Push(3)
	fmt.Println(stack.GetMin(),stack.GetMax(),stack.Peek())
	stack.Push(9)
	fmt.Println(stack.GetMin(),stack.GetMax(),stack.Peek())
	stack.Pop()
	fmt.Println(stack.GetMin(),stack.GetMax(),stack.Peek())


}