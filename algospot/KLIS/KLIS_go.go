package main

import (
	"fmt"
	"math"
	"sort"
)

func minInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ================ 1.2 Compatibility ===============
// For golang < 1.8 (use sort.Slice for 1.8 or above)
type item struct {
	Index int
	Value int
}
type itemSlice []item

func (a itemSlice) Len() int {
	return len(a)
}

func (a itemSlice) Less(i, j int) bool {
	return a[i].Value < a[j].Value
}

func (a itemSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// ==================================================

// Returns k-th LIS
func klis(seq []int, k int) []int {
	n := len(seq)

	// Initialize caches
	lisCache := make([]int, n+1)
	clisCache := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		lisCache[i] = -1
		clisCache[i] = -1
	}

	// Returns longest length of LIS
	var lis func(int) int
	lis = func(pos int) int {
		ret := &lisCache[pos+1]
		if *ret != -1 {
			return *ret
		}

		*ret = 1
		for next := pos + 1; next < n; next++ {
			if pos == -1 || seq[next] > seq[pos] {
				*ret = maxInt(*ret, 1+lis(next))
			}
		}
		return *ret
	}

	// Returns count of LIS
	var clis func(int) int
	clis = func(pos int) int {
		if lis(pos) == 1 {
			return 1
		}

		ret := &clisCache[pos+1]
		if *ret != -1 {
			return *ret
		}
		*ret = 0
		for next := pos + 1; next < n; next++ {
			if (pos == -1 || seq[next] > seq[pos]) && lis(pos) == lis(next)+1 {
				// *ret += clis(next) // Fails for judge; 0 <= K < (2^31 - 1)
				*ret = int(minInt64(int64(math.MaxInt32), int64(*ret)+int64(clis(next))))
			}
		}
		return *ret
	}

	// Get k-th LIS
	var solve func(int, int, *[]int)
	solve = func(pos, skip int, result *[]int) {
		if pos > -1 {
			*result = append(*result, seq[pos])
		}

		// cand := make([]int, 0, n)
		cand := make(itemSlice, 0, n)
		for next := pos + 1; next < n; next++ {
			if (pos == -1 || seq[next] > seq[pos]) && lis(pos) == lis(next)+1 {
				// cand = append(cand, next)
				cand = append(cand, item{Index: next, Value: seq[next]})
			}
		}
		// sort.Slice(cand, func(i, j int) bool { return seq[cand[i]] < seq[cand[j]] })
		sort.Sort(cand)
		for _, next := range cand {
			cnt := clis(next.Index)
			if cnt <= skip {
				skip -= cnt
			} else {
				solve(next.Index, skip, result)
				break
			}
		}
	}

	result := new([]int)
	solve(-1, k-1, result)
	return *result
}

func main() {
	var C, N, K int
	fmt.Scan(&C)
	for ; C > 0; C-- {
		fmt.Scan(&N, &K)
		seq := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&seq[i])
		}
		result := klis(seq, K)
		rLen := len(result)
		fmt.Println(rLen)
		for i := 0; i < rLen; i++ { // Can use strings.Join with some modifications
			fmt.Print(result[i])
			if i < (rLen - 1) {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
