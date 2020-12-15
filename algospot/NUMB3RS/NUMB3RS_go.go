package main

import (
	"bytes"
	"fmt"
	"strings"
)

// Type alias for matrix
type matrix []row
type row []float64

// Returns power of matrix
// NOTE: no dimension checking implemeneted
func matpow(a matrix, n int) matrix {
	if n == 1 {
		return a
	}

	if n%2 == 1 {
		return matmul(a, matpow(a, n-1))
	}

	temp := matpow(a, n/2)
	return matmul(temp, temp)
}

// Returns multiplication of two matrix
// NOTE: no dimension checking implemeneted
func matmul(a, b matrix) matrix {
	n := len(a)
	c := make(matrix, n)
	for i := 0; i < n; i++ {
		c[i] = make(row, n)
		for j := 0; j < n; j++ {
			temp := 0.0
			for k := 0; k < n; k++ {
				temp += a[i][k] * b[k][j]
			}
			c[i][j] = temp
		}
	}
	return c
}

// Returns the probabilities of jailbreaker presence for each given :towns
func numb3rs(m matrix, days, prison int, query []int) []float64 {
	// Convert adjcency matrix into probability matrix
	// by dividing each field by count of 1
	for i := 0; i < len(m); i++ {
		// Get count of 1 for row
		count := 0
		for j := 0; j < len(m); j++ {
			if m[i][j] == 1 {
				count++
			}
		}
		// Divide
		for j := 0; j < len(m); j++ {
			m[i][j] /= float64(count)
		}
	}

	// Calculate m^days
	m = matpow(m, days)

	// Collect probabilities for query
	ret := make([]float64, 0, len(query))
	for _, q := range query {
		ret = append(ret, m[prison][q])
	}

	return ret
}

func main() {
	var c, n, d, p, t int
	fmt.Scan(&c)
	for ; c > 0; c-- {
		fmt.Scan(&n, &d, &p)
		matrix := make(matrix, n)
		for i := 0; i < n; i++ {
			matrix[i] = make(row, n)
			for j := 0; j < n; j++ {
				// Take input as float; see type :matrix
				fmt.Scan(&matrix[i][j])
			}
		}
		fmt.Scan(&t)
		towns := make([]int, t)
		for i := 0; i < t; i++ {
			fmt.Scan(&towns[i])
		}

		result := numb3rs(matrix, d, p, towns)

		// Process output
		var b bytes.Buffer
		for _, prob := range result {
			b.WriteString(fmt.Sprintf("%.8f ", prob))
		}
		str := strings.TrimSpace(b.String())
		fmt.Println(str)
	}
}
