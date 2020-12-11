package main

import "fmt"

var cache [501]int

// Returns bigger one
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Returns length of increasing subsequence
func lis(seq []int, start int) int {
	n := len(seq)

	// Out of range
	if start == n {
		return 0
	}

	// Retrieve cache
	cacheIdx := start + 1
	ret := cache[cacheIdx]
	if ret != -1 {
		return ret
	}
	ret = 0

	for next := start + 1; next < n; next++ {
		if start == -1 || seq[next] > seq[start] {
			ret = max(ret, 1+lis(seq, next))
		}
	}

	cache[cacheIdx] = ret
	return ret
}

func main() {
	var C, N int
	fmt.Scan(&C)
	for ; C > 0; C-- {
		fmt.Scan(&N)
		seq := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&seq[i])
		}

		// Initialize cache
		for i := 0; i < (N + 1); i++ {
			cache[i] = -1
		}

		result := lis(seq, -1)
		fmt.Println(result)
	}
}
