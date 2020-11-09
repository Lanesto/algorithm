package main

import (
	"fmt"
	"errors"
)

var ErrNodeAlreadyExist = errors.New("Node already exist in list")
var ErrNodeDoesNotExist = errors.New("Node does not exists")

type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

const MIN_INT = -int(^uint(0) >> 1) - 1

// Returns new HEAD pointer after inserting element
func addNode(t *Node, v int) (*Node, error) {
	if t == nil {
		return &Node{v, nil, nil}, nil
	}

	temp := &Node{MIN_INT, nil, t}
	t = temp
	for t.Next != nil {
		if t.Next.Value == v {
			return temp.Next, ErrNodeAlreadyExist
		} else if t.Next.Value > v{
			var inserted *Node
			if t == temp {
				inserted = &Node{v, nil, t.Next}
			} else {
				inserted = &Node{v, t, t.Next}
			}
			t.Next.Prev = inserted
			t.Next = inserted
			break
		}
		t = t.Next
	}

	if t.Next == nil {
		t.Next = &Node{v, t, nil}
	}

	return temp.Next, nil
}

func deleteNode(t *Node, v int) (*Node, error) {
	if t == nil {
		return nil, ErrNodeDoesNotExist
	}

	head := t
	for t != nil {
		if t.Value == v {
			if t == head {
				t.Next.Prev = nil
				return t.Next, nil
			}
			t.Next.Prev = t.Prev
			t.Prev.Next = t.Next
			return head, nil
		}
		t = t.Next
	}

	return head, ErrNodeDoesNotExist
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

func reverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	for t.Next != nil{
		t = t.Next
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Prev
	}
	fmt.Println()
}

func size(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}

	n := 0
	for t != nil {
		n++
		t = t.Next
	}
	return n
}

func lookupNode(t *Node, v int) bool {
	if t == nil {
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}

func main() {
	var root *Node
	var err error
	fmt.Println(root)
	traverse(root)
	if root, err = addNode(root, 1); err != nil { fmt.Println(err) }
	if root, err = addNode(root, 1); err != nil { fmt.Println(err) }
	traverse(root)
	root, _ = addNode(root, 10)
	root, _ = addNode(root, 5)
	root, _ = addNode(root, 0)
	root, _ = addNode(root, 0)
	root, _ = addNode(root, 100)
	root, _ = addNode(root, -3)
	traverse(root)
	root, _ = deleteNode(root, 0)
	root, _ = deleteNode(root, 0)
	fmt.Println("Size:", size(root))
	traverse(root)
	reverse(root)
}
