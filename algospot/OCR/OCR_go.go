package main

import (
	"fmt"
	"math"
	"strings"
)

const (
	cacheNotSet = 1.0
)

var (
	// +inf
	infinity = math.Inf(0)

	// For DP
	cache [102][502]float64

	// To reconstruct result
	choices [102][502]int
)

// indexOf returns index of element in given slice, -1 if not found.
func indexOf(elem string, data []string) int {
	for i, v := range data {
		if v == elem {
			return i
		}
	}
	return -1
}

// ocr returns the sentence Q maximizing P(Q|R), where R is what OCR recognized as.
func ocr(words []string, after, recognize [][]float64, query []string) string {
	// Prepend dummy string that NEVER be shown for during runtime.
	words = append([]string{"## ERROR ##"}, words...)

	// Convert :query into slice of index of word in :words
	R := make([]int, len(query))
	for i := 0; i < len(R); i++ {
		R[i] = indexOf(query[i], words)
	}

	// Initialize cache
	for i := 0; i < len(R)+1; i++ {
		for j := 0; j < len(words)+1; j++ {
			cache[i][j] = cacheNotSet
			// No need to initialize choices
			// choices[i][j] = -1
		}
	}

	// Calculate P(Q|R)
	var solve func(int, int) float64
	solve = func(segment, previous int) float64 {
		if segment == len(query) {
			return 0.0 // Log10(1)
		}

		ret := &cache[segment][previous]
		if *ret != cacheNotSet {
			return *ret
		}

		*ret = -infinity
		for next := 1; next < len(words); next++ {
			temp := after[previous][next] + recognize[next][R[segment]] + solve(segment+1, next)
			if temp > *ret {
				*ret = temp
				choices[segment][previous] = next
			}
		}
		return *ret
	}

	// Returns Q
	var reconstruct func() string
	reconstruct = func() string {
		previous := 0
		ret := ""
		for segment := 0; segment < len(R); segment++ {
			next := choices[segment][previous]
			ret += " " + words[next]
			previous = next
		}
		return strings.TrimSpace(ret)
	}

	solve(0, 0)
	Q := reconstruct()
	return Q
}

func main() {
	var m, q, l int
	var temp float64
	fmt.Scan(&m, &q)

	words := make([]string, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&words[i])
	}

	B := make([]float64, m+1)
	B[0] = -infinity
	for i := 1; i < m+1; i++ {
		fmt.Scan(&temp)
		B[i] = math.Log10(temp)
	}

	T := make([][]float64, m+1)
	T[0] = B
	for i := 1; i < m+1; i++ {
		T[i] = make([]float64, m+1)
		T[i][0] = -infinity // log_(0)
		for j := 1; j < m+1; j++ {
			fmt.Scan(&temp)
			T[i][j] = math.Log10(temp)
		}
	}

	M := make([][]float64, m+1)
	M[0] = make([]float64, m+1)
	for i := 0; i < m+1; i++ {
		M[0][i] = -infinity // log_(0)
	}
	for i := 1; i < m+1; i++ {
		M[i] = make([]float64, m+1)
		for j := 1; j < m+1; j++ {
			fmt.Scan(&temp)
			M[i][j] = math.Log10(temp)
		}
	}

	for i := 0; i < q; i++ {
		fmt.Scan(&l)
		query := make([]string, l)
		for j := 0; j < l; j++ {
			fmt.Scan(&query[j])
		}
		fmt.Println(ocr(words, T, M, query))
	}
}
