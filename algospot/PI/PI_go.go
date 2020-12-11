package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var (
	// Cache func score(); optional.
	scoreCache = make(map[string]int, 100000)
	// Cache func pi()
	piCache [10000]int
)

// Calculate the difficulty of string
func score(pisubstr string) int {
	if v, ok := scoreCache[pisubstr]; ok {
		return v
	}

	var inner func(string) int
	inner = func(pisubstr string) int {
		n := len(pisubstr)
		if n < 3 || n > 5 {
			return math.MaxInt32
		}

		// Calculate differences of adjcent elements
		diffs := make([]int, n-1)
		for i := 0; i < len(diffs); i++ {
			diffs[i] = int(pisubstr[i+1]) - int(pisubstr[i])
		}

		// Evaluate properties
		prev := diffs[0]
		equalDiff, alternating := true, true
		for i := 1; i < len(diffs); i++ {
			if diffs[i] != prev {
				equalDiff = false
			}
			if diffs[i] != -prev {
				alternating = false
			}
			if !equalDiff && !alternating {
				break
			}
			prev = diffs[i]
		}

		// Judge
		if equalDiff {
			if diffs[0] == 0 {
				return 1
			} else if diffs[0] == -1 || diffs[0] == 1 {
				return 2
			}
			return 5
		} else if alternating {
			return 4
		}
		return 10
	}

	temp := inner(pisubstr)
	scoreCache[pisubstr] = temp
	return temp
}

func pi(pistr string) int {
	n := len(pistr)

	var solve func(int) int
	solve = func(start int) int {
		if start >= n {
			return 0
		}

		// Lookup cache
		ret := piCache[start]
		if ret != -1 {
			return ret
		}

		ret = math.MaxInt32
		for size := 3; size <= 5; size++ {
			end := min(start+size, n)
			pisubstr := pistr[start:end]
			ret = min(ret, score(pisubstr)+solve(end))
		}
		piCache[start] = ret
		return ret
	}

	return solve(0)
}

func main() {
	var C int
	var pistr string
	fmt.Scan(&C)
	for ; C > 0; C-- {
		fmt.Scan(&pistr)

		// Initialize cache
		for i := 0; i < len(pistr)+1; i++ {
			piCache[i] = -1
		}

		result := pi(pistr)
		fmt.Println(result)
	}
}
