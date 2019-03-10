package helperFunc

type minMaxStack struct {
	st       []int
	minStack []int
	maxStack []int
}

func NewMinMaxStack() *minMaxStack {
	minMaxStack := minMaxStack{st: make([]int, 0), minStack: make([]int, 0), maxStack: make([]int, 0)}
	return &minMaxStack

}

func (stack *minMaxStack) Peek() int {
	return stack.st[len(stack.st)-1]
}

func (stack *minMaxStack) Pop() int {
	poppedElement := stack.st[len(stack.st)-1]
	stack.st = stack.st[:len(stack.st)-1]
	stack.minStack = stack.minStack[:len(stack.minStack)-1]
	stack.maxStack = stack.maxStack[:len(stack.maxStack)-1]

	return poppedElement

}

func (stack *minMaxStack) Push(number int) int {
	stack.st = append(stack.st, number)

	if len(stack.minStack) == 0 {
		stack.minStack = append(stack.minStack, number)
		stack.maxStack = append(stack.maxStack, number)
		return stack.st[len(stack.st)-1]
	}

	if stack.minStack[len(stack.minStack)-1] > number {
		stack.minStack = append(stack.minStack, number)
	} else {
		stack.minStack = append(stack.minStack, stack.minStack[len(stack.minStack)-1])
	}

	if stack.maxStack[len(stack.maxStack)-1] < number {
		stack.maxStack = append(stack.maxStack, number)
	} else {
		stack.maxStack = append(stack.maxStack, stack.maxStack[len(stack.maxStack)-1])
	}

	return stack.st[len(stack.st)-1]

}

func (stack *minMaxStack) GetMin() int {
	return stack.minStack[len(stack.minStack)-1]
}

func (stack *minMaxStack) GetMax() int {
	return stack.maxStack[len(stack.maxStack)-1]
}
