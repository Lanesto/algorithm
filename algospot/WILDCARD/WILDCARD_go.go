package main

import (
	"fmt"
	"sort"
	"sync"
)

// Returns whether filename matches to a given pattern
// It implements caching internally
func solve(pattern, filename string) bool {
	// For cache
	const (
		Unknown = iota
		True
		False
	)

	// Initialize the cache and shared variables across recursive run
	plen, flen := len(pattern), len(filename)
	pn, fn := plen+1, flen+1
	cache := make([][]int, pn)
	for i := 0; i < pn; i++ {
		cache[i] = make([]int, fn)
		for j := 0; j < fn; j++ {
			cache[i][j] = Unknown
		}
	}

	// Runner
	var run func(int, int) bool
	run = func(pp, fp int) bool {
		// Consumed all characters in pattern or filename space
		if pp == plen {
			return fp == flen
		}
		// Lookup cache
		c := &cache[pp][fp]
		if *c != Unknown {
			return *c == True
		}
		// Match character or ?; advance 1 for each space
		*c = False
		if pp < plen && fp < flen && (filename[fp] == pattern[pp] || pattern[pp] == '?') {
			temp := run(pp+1, fp+1)
			if temp {
				*c = True
			}
			return temp
		}
		// Wildcard character; consume 1 pattern and consume 0~all remaining filename characters
		if pattern[pp] == '*' {
			for i := fp; i <= flen; i++ {
				if run(pp+1, i) {
					*c = True
					return true
				}
			}
		}
		*c = False
		return false
	}

	return run(0, 0)
}

// Filter filenames that match to patten
// Run multiple goroutines for every files.
func filterMatches(pattern string, filenames []string) []string {
	ret := make([]string, 0, len(filenames))
	var w sync.WaitGroup
	var m sync.Mutex
	for _, filename := range filenames {
		w.Add(1)
		// Concurrent run
		go func(filename string) {
			defer w.Done()
			if solve(pattern, filename) {
				m.Lock()
				ret = append(ret, filename)
				m.Unlock()
			}
		}(filename)
	}
	w.Wait()
	sort.Strings(ret)
	return ret
}

func main() {
	var C, N int
	var W string
	fmt.Scan(&C)
	for ; C > 0; C-- {
		fmt.Scan(&W, &N)
		filenames := make([]string, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&filenames[i])
		}
		matches := filterMatches(W, filenames)
		for _, match := range matches {
			fmt.Println(match)
		}
	}
}
