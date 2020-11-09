package main

import "fmt"

type float float64

type Node struct {
	Value float
	Next  *Node
}

var size = 0
var queue = new(Node)

func Push(t *Node, v float) bool {
	if queue == nil {
		queue = &Node{v, nil}
		size++
		return true
	}

	t = &Node{v, nil}
	t.Next = queue
	queue = t

	size++

	return true
}

func Pop(t *Node) (float, bool) {
	if size == 0 {
		return 0, false
	}

	if size == 1 {
		queue = nil
		size--
		return t.Value, true
	}

	tmp := t
	for t.Next != nil {
		tmp = t
		t = t.Next
	}

	v := tmp.Next.Value
	tmp.Next = nil

	size--
	return v, true
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("Empty Queue!")
		return
	}

	for t != nil {
		fmt.Printf("%f -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func main() {
	queue = nil
	Push(queue, 10)
	fmt.Println("Size:", size)
	traverse(queue)

	v, b := Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)

	for i := 0; i < 5; i++ {
		Push(queue, float(i))
	}
	traverse(queue)
	fmt.Println("Size:", size)

	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)

	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)
	traverse(queue)
}
