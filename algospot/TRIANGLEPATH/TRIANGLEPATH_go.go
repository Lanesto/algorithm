package main

import "fmt"

var cache [101][101]int

// Compare a and b then returns bigger one
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Returns the maximum path
func tripath(triangle [][]int, x, y int) int {
	if y == len(triangle) {
		return 0
	}
	ret := cache[y][x]
	if ret != -1 {
		return ret
	}
	ret = triangle[y][x] + max(tripath(triangle, x, y+1), tripath(triangle, x+1, y+1))
	cache[y][x] = ret
	return ret
}

func main() {
	var C, n int
	fmt.Scan(&C)
	for ; C > 0; C-- {
		fmt.Scan(&n)
		triangle := make([][]int, n)
		for i := 0; i < n; i++ {
			triangle[i] = make([]int, (i + 1))
			for j := 0; j < (i + 1); j++ {
				fmt.Scan(&triangle[i][j])
			}
		}

		// Initialize cache
		for i := 0; i < (n + 1); i++ {
			for j := 0; j < (n + 1); j++ {
				cache[i][j] = -1
			}
		}

		result := tripath(triangle, 0, 0)
		fmt.Println(result)
	}
}
