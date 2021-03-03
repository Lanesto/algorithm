package main

import (
	"fmt"
	"strings"
)

func dragon(gen, start, length int) string {
	// Convert human-friendly index (1...) to computer-friendly (0...)
	start--

	// Calculcate length of symbol after gen generation
	// X => X+YF (X, Y, +, F); xln[n] = xln[n-1] + yln[n-1] + 2
	// Y => FX-Y (X, Y, -, F); yln[n] = yln[n-1] + xln[n-1] + 2
	// ∴ xln[n] = 2 * xln[n-1] + 2 (∵ xln[n] == yln[n])
	xln := make([]int, gen+1)
	xln[0] = 1
	for i := 1; i < gen+1; i++ {
		xln[i] = 2 * (xln[i-1] + 1)
	}

	// Get k-th character of n-generation dragon curve
	var get func(string, int, int) string
	get = func(curve string, n, skip int) string {
		if n == 0 { // Where expanded symbols are picked
			return string(curve[skip])
		}

		for _, char := range curve {
			if char == 'X' || char == 'Y' { // Expand symbol
				// Whatever current value of skip is, symbol X and Y must be expanded
				// Because its first character changes every generations.
				cnt := xln[n]
				if cnt <= skip {
					skip -= cnt
				} else {
					if char == 'X' {
						return get("X+YF", n-1, skip)
					}
					return get("FX-Y", n-1, skip)
				}
			} else if skip == 0 {
				return string(char)
			} else {
				// Non-expanding symbol
				skip--
			}
		}
		// Never be returned
		return "#"
	}

	tmp := make([]string, length)
	for i := 0; i < length; i++ {
		char := get("FX", gen, start+i)
		tmp = append(tmp, char)
	}

	return strings.Join(tmp, "")
}

func main() {
	var c, n, p, l int
	fmt.Scan(&c)
	for ; c > 0; c-- {
		fmt.Scan(&n, &p, &l)
		result := dragon(n, p, l)
		fmt.Println(result)
	}
}
