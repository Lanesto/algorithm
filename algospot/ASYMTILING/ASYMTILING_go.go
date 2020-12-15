package main

import "fmt"

var (
	tilingCache     [101]int
	asymtilingCache [101]int
)

// Same as TILING2
func tiling(n int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	}

	// Retrieve cache
	ret := &tilingCache[n]
	if *ret != -1 {
		return *ret
	}

	*ret = (tiling(n-1) + tiling(n-2)) % 1000000007
	return *ret
}

// Returns number of cases tiling but is asymmetric
func asymtiling(n int) int {
	if n <= 2 {
		return 0
	}

	// Retrieve cache
	ret := &asymtilingCache[n]
	if *ret != -1 {
		return *ret
	}

	*ret = (asymtiling(n-4) + asymtiling(n-2) + 2*(tiling(n-3))) % 1000000007
	return *ret
}

func init() {
	for i := 0; i < 101; i++ {
		tilingCache[i] = -1
		asymtilingCache[i] = -1
	}
}

func main() {
	var C, n int
	fmt.Scan(&C)
	for ; C > 0; C-- {
		fmt.Scan(&n)
		result := asymtiling(n)
		fmt.Println(result)
	}
}
