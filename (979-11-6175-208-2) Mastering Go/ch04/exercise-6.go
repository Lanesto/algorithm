package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func sqrt(f float64, precision uint) *big.Float {
	k := 0
	result := new(big.Float).SetPrec(precision).SetFloat64(float64(0.5) * (float64(1.0) + (float64(f) / float64(1.0))))
	for {
		if k > int(precision) {
			break
		}
		temp := new(big.Float).SetFloat64(f) 
		temp.Quo(temp, result)
		result.Add(result, temp)
		result.Mul(result, new(big.Float).SetFloat64(0.5))
		k++
	}
	return result
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Insufficient arguments; requires number, precision")
		os.Exit(1)
	}

	f, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("number must be floating point number")
		os.Exit(1)
	}

	prec, err := strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil {
		fmt.Println("precision must be integer")
		os.Exit(1)
	}

	fmt.Println(sqrt(f, uint(prec)))
}
