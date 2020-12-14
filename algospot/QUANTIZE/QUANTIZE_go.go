package main

import (
	"fmt"
	"math"
	"sort"
)

// Used to cache results in quantize(), internally
// must be initialized every time quantize() is called
// cache[0 <= start <= 100][0 <= k <= 10]
var cache [101][11]int

// Returns smaller one
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Returns round of float as integer
// math.Round is available >= 1.10
func round(f float64) int {
	t := math.Trunc(f)
	if math.Abs(f-t) >= 0.5 {
		return int(t) + 1
	}
	return int(t)
}

// Quantize given sequence into k numbers minimizing
// the sum of SSE, Sum of Squared Error
func quantize(seq []int, k int) int {
	n := len(seq)

	// Sort the sequence
	sort.Ints(seq)

	// Initialize ps(Partial Sum), pss(Partial Sum of Squared)
	// ps[i] is sum of seq[0:i]; [0, i)
	// pss[i] is sum of squared seq[0:i]; [0, i)
	ps := make([]int, n+1)
	pss := make([]int, n+1)
	for i := 0; i < n; i++ {
		ps[i+1] = ps[i] + seq[i]
		pss[i+1] = pss[i] + (seq[i] * seq[i])
	}

	// Returns the SSE of given range, [l, r)
	sse := func(l, r int) int {
		s := ps[r] - ps[l]
		ss := pss[r] - pss[l]
		m := round(float64(s) / float64(r-l))
		return ss - 2*m*s + (r-l)*m*m
	}

	// Initialize cache
	for i := 0; i < n+1; i++ {
		for j := 0; j < k+1; j++ {
			cache[i][j] = -1
		}
	}

	var solve func(int, int) int
	solve = func(start, k int) int {
		// Base conditions
		if start == n {
			// Used all numbers
			return 0
		} else if k == 0 {
			// End with unused numbers
			return math.MaxInt32
		}

		// Retrieve cache
		ret := cache[start][k]
		if ret != -1 {
			return ret
		}

		ret = math.MaxInt32
		for size := 1; start+size <= n; size++ {
			ret = min(ret, sse(start, start+size)+solve(start+size, k-1))
		}
		cache[start][k] = ret
		return ret
	}

	return solve(0, k)
}

func main() {
	var C, N, S int
	fmt.Scan(&C)
	for ; C > 0; C-- {
		fmt.Scan(&N, &S)
		seq := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&seq[i])
		}
		result := quantize(seq, S)
		fmt.Println(result)
	}
}
