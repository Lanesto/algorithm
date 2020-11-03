package main

import "fmt"

const (
	sq_4_0 = 1 << (iota * 2)
	sq_4_1
	sq_4_2
	sq_4_3
	sq_4_4
	sq_4_5
)

const (
	Mon = iota
	Tue
	Wed
	Thu
	Fri
	Sat
	Sun
)

func main() {
	fmt.Println("4^0:", sq_4_0)
	fmt.Println("4^1:", sq_4_1)
	fmt.Println("4^2:", sq_4_2)
	fmt.Println("4^3:", sq_4_3)
	fmt.Println("4^4:", sq_4_4)
	fmt.Println("4^5:", sq_4_5)

	fmt.Println("Mon:", Mon)
	fmt.Println("Tue:", Tue)
	fmt.Println("Wed:", Wed)
	fmt.Println("Thu:", Thu)
	fmt.Println("Fri:", Fri)
	fmt.Println("Sat:", Sat)
	fmt.Println("Sun:", Sun)
}
