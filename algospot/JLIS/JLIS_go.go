package main

import (
	"fmt"
	"math"
)

var cache [101][101]int

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func jlis(seqA, seqB []int, startA, startB int) int {
	// Select base number
	var valA int64 = math.MinInt64
	if startA >= 0 {
		valA = int64(seqA[startA])
	}
	var valB int64 = math.MinInt64
	if startB >= 0 {
		valB = int64(seqB[startB])
	}
	base := maxInt64(valA, valB)

	// Lookup cache
	ret := cache[startA+1][startB+1]
	if ret != -1 {
		return ret
	}
	ret = 0

	// Try other numbers with fixed B
	for nextA := startA + 1; nextA < len(seqA); nextA++ {
		if int64(seqA[nextA]) > base {
			ret = maxInt(ret, 1+jlis(seqA, seqB, nextA, startB))
		}
	}

	// Try other numbers with fixed A
	for nextB := startB + 1; nextB < len(seqB); nextB++ {
		if int64(seqB[nextB]) > base {
			ret = maxInt(ret, 1+jlis(seqA, seqB, startA, nextB))
		}
	}

	cache[startA+1][startB+1] = ret
	return ret
}

func main() {
	var c, n, m int
	fmt.Scan(&c)
	for ; c > 0; c-- {
		fmt.Scan(&n, &m)
		seqA := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&seqA[i])
		}
		seqB := make([]int, m)
		for i := 0; i < m; i++ {
			fmt.Scan(&seqB[i])
		}

		for i := 0; i < (n + 1); i++ {
			for j := 0; j < (m + 1); j++ {
				cache[i][j] = -1
			}
		}

		result := jlis(seqA, seqB, -1, -1)
		fmt.Println(result)
	}
}
