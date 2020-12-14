package main

import "fmt"

var cache [101]int

// Returns the number of case covering 2*n sqaure with 2*1 tiles
func tiling2(n int) int {
	if n <= 1 {
		return 1
	}

	ret := cache[n]
	if ret != -1 {
		return ret
	}

	ret = (tiling2(n-1)%1000000007 + tiling2(n-2)%1000000007) % 1000000007
	cache[n] = ret
	return ret
}

func init() {
	for i := 0; i < 101; i++ {
		cache[i] = -1
	}
}

func main() {
	var C, n int
	fmt.Scan(&C)
	for ; C > 0; C-- {
		fmt.Scan(&n)
		result := tiling2(n)
		fmt.Println(result)
	}
}
