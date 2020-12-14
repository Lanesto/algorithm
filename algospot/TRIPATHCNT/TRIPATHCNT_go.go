package main

import "fmt"

var (
	pathCache [101][101]int
	cntCache  [101][101]int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func tripathcnt(triangle [][]int) int {
	n := len(triangle)

	for i := 0; i < n+1; i++ {
		for j := 0; j < n+1; j++ {
			pathCache[i][j] = -1
			cntCache[i][j] = -1
		}
	}

	var tripath func(int, int) int
	tripath = func(x, y int) int {
		if y == n {
			return 0
		}

		ret := pathCache[y][x]
		if ret != -1 {
			return ret
		}

		ret = triangle[y][x] + max(tripath(x, y+1), tripath(x+1, y+1))
		pathCache[y][x] = ret
		return ret
	}

	var solve func(int, int) int
	solve = func(x, y int) int {
		if y == n-1 {
			return 1
		}

		ret := cntCache[y][x]
		if ret != -1 {
			return ret
		}

		ret = 0
		next := tripath(x, y) - triangle[y][x]
		if next == tripath(x, y+1) {
			ret += solve(x, y+1)
		}
		if next == tripath(x+1, y+1) {
			ret += solve(x+1, y+1)
		}
		cntCache[y][x] = ret
		return ret
	}

	return solve(0, 0)
}

func main() {
	var C, n int
	fmt.Scan(&C)
	for ; C > 0; C-- {
		fmt.Scan(&n)
		triangle := make([][]int, n)
		for i := 0; i < n; i++ {
			triangle[i] = make([]int, n)
			for j := 0; j < i+1; j++ {
				fmt.Scan(&triangle[i][j])
			}
		}
		result := tripathcnt(triangle)
		fmt.Println(result)
	}
}
