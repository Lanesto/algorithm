package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// Length of overlapping substring between two string a, b
//
// e.g) a[:lenOverlap(a, b)] + b; geo + oji -> geoji (ge, oji)
func lenOverlap(a, b string) int {
	for i := 0; i < len(a); i++ {
		if strings.HasPrefix(b, a[i:]) {
			return i
		}
	}
	return len(a)
}

// Returns shortest string made of substrs, with overlapping duplicates removed
func restore(substrs []string) string {
	// Find substring of substring that will be excluded
	excluded := make(map[string]struct{}) // Set
	for _, a := range substrs {
		for _, b := range substrs {
			if a == b {
				continue
			}
			if strings.Contains(a, b) {
				excluded[b] = struct{}{}
			}
		}
	}
	newSubstrs := make([]string, 0, len(substrs))
	for _, s := range substrs {
		if _, ok := excluded[s]; !ok { // Add items not excluded only
			newSubstrs = append(newSubstrs, s)
		}
	}

	// Not necessary but for test
	sort.Strings(newSubstrs)

	// Replace original with dummy string in front; dummy string is for index -1 (-> 0)
	substrs = append([]string{"## DUMMY ##"}, newSubstrs...)
	n := len(substrs)
	allTaken := (1 << uint(n)) - 1

	// Precalc overlap[i][j]
	overlap := make([][]int, n)
	for i := 0; i < n; i++ {
		overlap[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 { // overlap[0][*] = 0
				overlap[i][j] = 0
			} else {
				overlap[i][j] = lenOverlap(substrs[i], substrs[j])
			}
		}
	}

	// Init cache
	cache := make([][]int, n)
	choices := make([][]int, n) // choices[prev][taken] = next
	for i := 0; i < n; i++ {
		cache[i] = make([]int, 1<<uint(n))
		choices[i] = make([]int, 1<<uint(n))
		for j := 0; j < (1 << uint(n)); j++ {
			cache[i][j] = -1
			choices[i][j] = -1
		}
	}

	// Returns the minimum length of result string
	var solve func(int, int) int
	solve = func(prev, taken int) int {
		if taken == allTaken {
			return len(substrs[prev])
		}

		// Retrieve from cache
		ret := &cache[prev][taken]
		if *ret != -1 {
			// Cache hit
			return *ret
		}

		// Cache miss
		*ret = math.MaxInt16
		for next := 1; next < n; next++ {
			if taken&(1<<uint(next)) == 0 {
				tmp := overlap[prev][next] + solve(next, taken|(1<<uint(next)))
				if tmp < *ret {
					*ret = tmp
					choices[prev][taken] = next
				}
			}
		}
		return *ret
	}

	// Make full string from choices
	var reconstruct func(int, int) string
	reconstruct = func(prev, taken int) string {
		if taken == allTaken {
			return substrs[prev]
		}

		next := choices[prev][taken]
		return substrs[prev][:overlap[prev][next]] + reconstruct(next, taken|(1<<uint(next)))
	}

	solve(0, 1)
	return reconstruct(0, 1)
}

func main() {
	var C, k int
	fmt.Scan(&C)
	for ; C > 0; C-- {
		fmt.Scan(&k)
		substrs := make([]string, k)
		for i := 0; i < k; i++ {
			fmt.Scan(&substrs[i])
		}
		result := restore(substrs)
		fmt.Println(result)
	}
}
