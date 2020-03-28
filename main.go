package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func printNode(n *Node) {
	fmt.Print("Value: ", n.Value)

	if n.Left != nil {
		fmt.Print(" Left: ", n.Left.Value)
	}

	if n.Right != nil {
		fmt.Print(" Right: ", n.Right.Value)
	}

	fmt.Println()
}

func main() {
	test := &Node{Value: 9}
	right := &Node{Value: 11}

	test.Right = right
	printNode(test)
}
