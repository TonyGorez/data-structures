package main

import "fmt"

// Node struct of node of linkedlist
type Node struct {
	value string
	next  *Node
	prev  *Node
}

// List truct of linked list
type List struct {
	head   *Node
	last   *Node
	length int
}

// First return the first node
func (list *List) First() *Node {
	return list.head
}

// Last return the last node
func (list *List) Last() *Node {
	return list.last
}

// Push add a value to linked list
func (list *List) Push(value string) {
	node := &Node{value: value}
	if list.head == nil {
		list.head = node
	} else {
		list.last.next = node
		node.prev = list.last
	}
	list.last = node
	list.length++

}

// Next return the next node
func (node *Node) Next() *Node {
	return node.next
}

// Prev return the prev ndoe
func (node *Node) Prev() *Node {
	return node.prev
}

func main() {
	myList := &List{}
	myList.Push("Zacary")
	myList.Push("Jmiaa")
	myList.Push("Tony")
	fmt.Println(myList.length)

	firstValue := myList.First()
	fmt.Println(firstValue.value, firstValue.next.value)
	fmt.Println(firstValue.Next().Next())

	lastValue := myList.Last()
	fmt.Println(lastValue.value, lastValue.prev.value)
	fmt.Println(lastValue.Prev().Prev())
}
