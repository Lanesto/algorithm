package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("No value given")
		os.Exit(1)
	}

	var sum float64 = 0.0
	for i := 1; i < len(os.Args); i++ {
		f, _ := strconv.ParseFloat(os.Args[i], 64)
		sum += f
	}
	avg := sum / float64(len(os.Args) - 1)
	fmt.Println(avg)
}
