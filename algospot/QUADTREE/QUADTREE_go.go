package main

import "fmt"

// Reverse given quadtree
// Each reverse() consumes certain amount of characters
// Using *string to share overall progress conveniently
func reverse(quadtree *string) string {
	// Advance step by 1
	head := string((*quadtree)[0])
	*quadtree = (*quadtree)[1:]

	// If not starts with "x" returns head
	if head == "b" || head == "w" {
		return string(head)
	}

	// x1234 -> x3412
	ul := reverse(quadtree) // up-left
	ur := reverse(quadtree) // up-right
	dl := reverse(quadtree) // down-left
	dr := reverse(quadtree) // down-right
	return "x" + dl + dr + ul + ur
}

func main() {
	var C int
	fmt.Scan(&C)
	for ; C > 0; C-- {
		var quadtree string
		fmt.Scan(&quadtree)

		result := reverse(&quadtree)
		fmt.Println(result)
	}
}
