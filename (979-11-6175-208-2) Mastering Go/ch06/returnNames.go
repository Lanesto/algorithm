package main

import (
	"fmt"
	"os"
	"strconv"
)

func namedMinMax(x, y int) (min, max int) {
	if x > y {
		min, max = y, x
	} else {
		min, max = x, y
	}
	return
}

func minMax(x, y int) (min, max int) {
	if x > y {
		min, max = y, x
	} else {
		min, max = x, y
	}
	return min, max
}

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Requires at least 2 arguments")
		return
	}

	a1, _ := strconv.Atoi(args[1])
	a2, _ := strconv.Atoi(args[2])

	fmt.Println(minMax(a1, a2))
	min, max := minMax(a1, a2)
	fmt.Println(min, max)

	fmt.Println(namedMinMax(a1, a2))
	min, max = minMax(a1, a2)
	fmt.Println(min, max)
}
