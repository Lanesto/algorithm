package main

import "fmt"

type Node struct {
	Value int
	Next *Node
}

const MIN_INT = -int(^uint(0) >> 1)-1

// Returns new HEAD pointer after inserting new element
func addNode(t *Node, v int) *Node {
	if t == nil {
		return &Node{v, nil}
	}

	temp := &Node{MIN_INT, t}
	t = temp
	for t.Next != nil {
		if t.Next.Value == v {
			fmt.Println("Node already exists:", v)
			break
		} else if t.Next.Value > v {
			t.Next = &Node{v, t.Next}
			break
		}
		t = t.Next
	}

	if t.Next == nil {
		t.Next = &Node{v, nil}
	}

	return temp.Next
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func lookupNode(t *Node, v int) bool {
	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}

func size(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}

	i := 0
	for t != nil {
		i++
		t = t.Next
	}
	return i
}

func main() {
	var root *Node
	root = addNode(root, 1)
	root = addNode(root, -1)
	traverse(root)
	root = addNode(root, 10)
	root = addNode(root, 5)
	root = addNode(root, 45)
	root = addNode(root, 5)
	root = addNode(root, 5)
	traverse(root)
	root = addNode(root, 100)
	traverse(root)

	if lookupNode(root, 100) {
		fmt.Println("Node exists!")
	} else {
		fmt.Println("Node does not exist!")
	}

	if lookupNode(root, -100) {
		fmt.Println("Node exists!")
	} else {
		fmt.Println("Node does not exist!")
	}
}



