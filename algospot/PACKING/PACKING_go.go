package main

import (
	"fmt"
)

func packing(capacity int, name []string, volume, pref []int) (int, []string) {
	n := len(name)

	// Init cache[capacity + 1][which + 1]
	cache := make([][]int, capacity+1)
	for i := 0; i < capacity+1; i++ {
		cache[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			cache[i][j] = -1
		}
	}

	// Solver
	var solve func(int, int) int
	solve = func(cap, which int) int {
		// Seen all items
		if which == n {
			return 0
		}

		ret := &cache[cap][which]
		if *ret != -1 {
			return *ret
		}

		*ret = solve(cap, which+1)
		if cap >= volume[which] {
			tmp := pref[which] + solve(cap-volume[which], which+1)
			if tmp > *ret {
				*ret = tmp
			}
		}
		return *ret
	}
	score := solve(capacity, 0)

	// Reconstruct
	picked := make([]string, 0)
	cap := capacity
	for which := 0; which < n; which++ {
		if solve(cap, which) == solve(cap, which+1) {
			// Not picked
		} else {
			// Picked
			picked = append(picked, name[which])
			cap -= volume[which]
		}
	}

	// sort.Strings(picked)
	return score, picked
}

func main() {
	var C, N, W int
	fmt.Scan(&C)
	for i := 0; i < C; i++ {
		fmt.Scan(&N, &W)
		var name = make([]string, N)
		var volume = make([]int, N)
		var pref = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&name[j], &volume[j], &pref[j])
		}
		sum, list := packing(W, name, volume, pref)
		fmt.Println(sum, len(list))
		for _, item := range list {
			fmt.Println(item)
		}
	}

}
