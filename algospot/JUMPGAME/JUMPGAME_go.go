package main

import "fmt"

const maxBoardSize = 100

var cache [maxBoardSize][maxBoardSize]int

// Returns whether is destination of board reachable.
// It includes caching behavior
func reachable(board [][]int, x, y int) bool {
	const (
		Unknown = iota
		True
		False
	)

	// Initialize cache
	boardSize := len(board)
	for i := 0; i < boardSize; i++ {
		// Set defaults explicitly
		// It is not necessary because make() fills 'zero' of type 'int';
		for j := 0; j < boardSize; j++ {
			cache[i][j] = Unknown
		}
	}

	// Solver; defined a variable for recursion
	var solve func([][]int, int, int) bool
	solve = func(board [][]int, x, y int) bool {
		// Out of board
		if (x < 0 || x >= boardSize) || (y < 0 || y >= boardSize) {
			return false
		}

		// Return cached value if exists
		if cache[y][x] != Unknown {
			return cache[y][x] == True
		}

		// Reached at destination
		if (x == boardSize-1) && (y == boardSize-1) {
			cache[y][x] = True
			return true
		}
		step := board[y][x]
		ret := solve(board, x+step, y) || solve(board, x, y+step)
		if ret {
			cache[y][x] = True
		} else {
			cache[y][x] = False
		}
		return ret
	}
	return solve(board, 0, 0)
}

func main() {
	// IO
	var C, n int
	fmt.Scan(&C)
	for ; C > 0; C-- {
		fmt.Scan(&n)
		board := make([][]int, n)
		for i := 0; i < n; i++ {
			row := make([]int, n)
			for j := 0; j < n; j++ {
				fmt.Scan(&row[j])
			}
			board[i] = row
		}
		// Convert and print output
		if reachable(board, 0, 0) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
