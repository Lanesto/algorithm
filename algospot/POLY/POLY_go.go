package main

import "fmt"

var cache [101][101]int

// Returns number of polynomioes
func poly(n int) int {

	// :k      number of blocks
	// :first  number of blocks to be placed at top among :k
	var solve func(int, int) int
	solve = func(k, first int) int {
		if k == 0 || k == first {
			return 1
		}

		// Retrieve cache
		ret := &cache[k][first]
		if *ret != -1 {
			return *ret
		}

		temp := 0
		for second := 1; second <= k-first; second++ {
			temp += (first + second - 1) * solve(k-first, second)
		}
		*ret = temp % 10000000
		return *ret
	}

	ret := 0
	for first := 1; first <= n; first++ {
		ret += solve(n, first)
	}
	return ret % 10000000
}

func init() {
	for i := 0; i < 101; i++ {
		for j := 0; j < 101; j++ {
			cache[i][j] = -1
		}
	}
}

func main() {
	var C, n int
	fmt.Scan(&C)
	for ; C > 0; C-- {
		fmt.Scan(&n)
		result := poly(n)
		fmt.Println(result)
	}
}
