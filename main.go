package main

import "fmt"

// Node is a vertex in a graph
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func read() []Node {
	var N int
	fmt.Scanf("%d", &N)

	nodes := make([]Node, N)
	fmt.Println(nodes)

	for i := 0; i < N; i++ {
		var val, indexL, indexR int
		fmt.Scanf("%d %d %d", &val, &indexL, &indexR)
		nodes[i].Value = val

		// remember these are just references to the nodes slice which was created up above with the make
		if indexL >= 0 {
			nodes[i].Left = &nodes[indexL]
		}

		if indexR >= 0 {
			nodes[i].Right = &nodes[indexR]
		}
	}

	return nodes
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
