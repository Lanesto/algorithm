package main

import (
	"fmt"
)

// Returns the maximum difference in scores of both players, in point of
// game starter player.
func numbergame(board []int) int {
	n := len(board)

	// Initialize cache
	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, n)
		for j := 0; j < n; j++ {
			cache[i][j] = -1
		}
	}

	// Subproblem for board in range [left, right], not [left, right)
	var solve func(int, int) int
	solve = func(left, right int) int {
		// Base problem: if there are only one number remaining, player must take it
		if left == right {
			return board[left] // No minus for it!
		}

		// Retrieve cache
		ret := &cache[left][right]
		if *ret != -1 {
			return *ret
		}

		// Remove two numbers from sides if there are enough numbers.
		if right-left >= 2 {
			*ret = maxInt(
				-solve(left+2, right),
				-solve(left, right-2),
			)
		}

		// Pick one from sides
		*ret = maxInt(
			*ret,
			board[left]-solve(left+1, right),
			board[right]-solve(left, right-1),
		)

		return *ret
	}
	result := solve(0, n-1)
	return result
}

func main() {
	var C, n int
	fmt.Scan(&C)
	for ; C > 0; C-- {
		fmt.Scan(&n)
		board := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&board[i])
		}
		result := numbergame(board)
		fmt.Println(result)
	}
}

// Returns the biggest integer from given integers.
func maxInt(nums ...int) int {
	max := nums[0]
	if len(nums) > 1 {
		for _, v := range nums[1:] {
			if v > max {
				max = v
			}
		}
	}
	return max
}
