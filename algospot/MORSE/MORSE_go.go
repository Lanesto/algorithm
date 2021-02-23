package main

import (
	"fmt"
	"strings"
)

var (
	bino [][]int
)

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// Returns k-th morse signal with n "-" and m "o", in dictionary order
func morse(n, m, k int) string {
	var solve func(int, int, int) string
	solve = func(n, m, skip int) string {
		if n == 0 {
			return strings.Repeat("o", m)
		}

		mid := bino[n+m-1][n-1] // Expected to be <= 10**9
		if skip < mid {
			return "-" + solve(n-1, m, skip)
		}
		return "o" + solve(n, m-1, skip-mid)
	}
	return solve(n, m, k-1)
}

func init() {
	max := 1000000100
	bino = make([][]int, 201)
	for i := 0; i < 200; i++ { // 1 <= n, m <= 100
		bino[i] = make([]int, 201)
		bino[i][0] = 1
		bino[i][i] = 1
		for j := 1; j < i; j++ {
			bino[i][j] = minInt(max, bino[i-1][j-1]+bino[i-1][j])
		}
	}
}

func main() {
	var C, n, m, k int
	fmt.Scan(&C)
	for ; C > 0; C-- {
		fmt.Scan(&n, &m, &k)
		result := morse(n, m, k)
		fmt.Println(result)
	}
}
