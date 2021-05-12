package main

import (
	"fmt"
	"sort"
)

// Returns maximum win count.
func matchOrder(russia, korea []int) int {
	sort.Ints(korea)
	strongest := korea[len(korea)-1]
	win := 0

	for _, rating := range russia {
		if rating > strongest {
			korea = korea[1:] // Removes first element
		} else {
			// Removes first element that is bigger than rating.
			idx := sort.SearchInts(korea, rating)
			korea = append(korea[:idx], korea[idx+1:]...)
			win++
		}
	}
	return win
}

func main() {
	var C, N int
	fmt.Scan(&C)
	for ; C > 0; C-- {
		fmt.Scan(&N)

		russia, korea := make([]int, N), make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&russia[i])
		}
		for j := 0; j < N; j++ {
			fmt.Scan(&korea[j])
		}

		result := matchOrder(russia, korea)
		fmt.Println(result)
	}
}
