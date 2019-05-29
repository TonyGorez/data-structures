package main

import (
	"fmt"
	"log"
)

type Node struct {
	value int
	left *Node
	right *Node
}

func printNode(n* Node) {
	fmt.Print("Value : ", n.value)
	if n.left != nil {
		fmt.Print(", Left : ", n.left.value)
	} else {
		fmt.Print(" No left value")
	}
	if n.right != nil {
		fmt.Print(", Right : ", n.right.value)
	} else {
		fmt.Print(" No right value")
	}
	fmt.Println()
}

func read() []Node {
	var N int
	_, err := fmt.Scanf("%d", &N)
	if err != nil {
		log.Fatal(err)
	}
	nodes := make([]Node, N)
	for i := 0; i < N; i++ {
		var val, indexLeft, indexRight int

		_, err := fmt.Scanf("%d %d %d", &val, &indexLeft, &indexRight)
		if err != nil {
			log.Fatal(err)
		}

		nodes[i].value = val
		if indexLeft >= 0 {
			nodes[i].left = &nodes[indexLeft]
		}
		if indexRight >= 0 {
			nodes[i].right = &nodes[indexRight]
		}
	}
	return nodes
}

func main() {
	nodes := read()
	for _, node := range nodes {
		printNode(&node)
	}
}
