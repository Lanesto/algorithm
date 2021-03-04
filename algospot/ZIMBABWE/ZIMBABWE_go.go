package main

import (
	"fmt"
	"sort"
)

// Convert byte representating 0 ~ 9 (48 ~ 57) to type int
func byteToInt(b byte) int {
	return int(b - 48)
}

// true to 1, false to 0
func boolToInt(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

// Returns the count of possible previous prices of egg.
func zimbabwe(eggPrice string, multiple int) int {
	// Preprocess
	n := len(eggPrice)
	digits := make([]int, n)
	for i, b := range []byte(eggPrice) {
		digits[i] = byteToInt(b) // or int(b - '0') when loop over runes
	}
	sort.Ints(digits)

	// Init cache
	cache := make([][][]int, (1 << uint(n))) // cache[taken][mod][less]
	for i := 0; i < (1 << uint(n)); i++ {
		cache[i] = make([][]int, 20+1)
		for j := 0; j < 20+1; j++ {
			cache[i][j] = make([]int, 2)
			for k := 0; k < 2; k++ {
				cache[i][j][k] = -1
			}
		}
	}

	// Solver
	var solve func(int, int, int, bool) int
	solve = func(
		pos int, // Index of current digit of eggPrice
		taken int, // Bitmask for each digits whether used or not
		mod int, // Partial modulo; abc % m = 100a % m + 10b % m + c % m
		less bool, // Flag whether partial price candidate are less than eggPrice
	) (cnt int) {
		// Base condition: visited all digits
		if pos == n { // or taken == (1 << n) - 1
			if less && mod == 0 {
				return 1
			}
			return 0
		}

		// Retrieve cache
		ret := &cache[taken][mod][boolToInt(less)]
		if *ret != -1 {
			return *ret
		}

		*ret = 0
		for i := 0; i < n; i++ {
			// 1 << uint(i) ( < 1.13)
			// 1 << i       (>= 1.13)
			// See proposal https://github.com/golang/proposal/blob/master/design/19113-signed-shift-counts.md
			if taken&(1<<uint(i)) == 0 {
				if !less && digits[i] > byteToInt(eggPrice[pos]) {
					// If eggPrice starts with substring already made, following digits
					// must be equal or smaller to digit of current position in eggPrice.
					continue
				}
				if i > 0 && digits[i-1] == digits[i] && taken&(1<<uint(i-1)) == 0 {
					// If there are same numbers in digits, force earliest one to be used
					// to remove duplicates.
					continue
				}
				*ret += solve(
					pos+1,
					taken|(1<<uint(i)),
					(10*mod+digits[i])%multiple,
					less || digits[i] < byteToInt(eggPrice[pos]),
				) % 1000000007
			}
		}
		return *ret
	}

	return solve(0, 0, 0, false) % 1000000007
}

func main() {
	var c, m int
	var e string
	fmt.Scan(&c)
	for ; c > 0; c-- {
		fmt.Scan(&e, &m)
		result := zimbabwe(e, m)
		fmt.Println(result)
	}
}
