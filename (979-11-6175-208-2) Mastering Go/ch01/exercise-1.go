package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	var sum int64 = 0
	for i := 1; i < len(os.Args); i++ {
		n, _ := strconv.ParseInt(os.Args[i], 10, 64)
		sum += n
	}
	fmt.Println(sum)
}
