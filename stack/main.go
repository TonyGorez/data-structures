package main

import "fmt"

// Stack type
type Stack struct {
	items  []string
	length int
}

// Push element to stack
func (stack *Stack) Push(element string) {
	stack.items = append(stack.items, element)
	stack.length++
}

// Pop element from stack
func (stack *Stack) Pop() {
	stack.items = stack.items[0 : len(stack.items)-1]
	stack.length--
}

func main() {

	myStack := Stack{}
	myStack.Push("Tony")
	myStack.Push("Jami")
	myStack.Push("Zac")

	fmt.Println("Befor", myStack)
	myStack.Pop()
	fmt.Println("After", myStack)
}
