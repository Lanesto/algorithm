package main

import "fmt"

var cache [1001][1001]float64

// Returns the robability to climb up :height in :days
func snail(height, days int) float64 {
	// Always succeed
	if height <= days {
		return 1.0
	}

	// Timeout
	if days <= 0 {
		if height > 0 {
			return 0.0
		}
		return 1.0
	}

	// Retrieve cache
	ret := &cache[height][days]
	if *ret != -0.5 {
		return *ret
	}

	*ret = 0.75*snail(height-2, days-1) + 0.25*snail(height-1, days-1)
	return *ret
}

func init() {
	for i := 0; i < 1001; i++ {
		for j := 0; j < 1001; j++ {
			cache[i][j] = -0.5
		}
	}
}

func main() {
	var C, n, m int
	fmt.Scan(&C)
	for ; C > 0; C-- {
		fmt.Scan(&n, &m)
		result := snail(n, m)
		fmt.Printf("%.10f\n", result)
	}
}
