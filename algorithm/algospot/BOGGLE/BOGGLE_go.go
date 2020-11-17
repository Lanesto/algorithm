package main

import "fmt"

// Size of boggle board (5x5)
const boardSizeX = 5
const boardSizeY = 5

// Maximum length of word to find given
const wordLengthMax = 10

// Caching for DP
var cache = [boardSizeX][boardSizeY][wordLengthMax]byte{}

const (
	none = iota
	no
	yes
)

// Possible directions
var dx = [8]int{-1, -1, -1,  0, 0,  1, 1, 1}
var dy = [8]int{-1,  0,  1, -1, 1, -1, 0, 1}

// From position (x, y) can find remaining subword (*word)[pos:]
func hasWord(board *[]string, x, y int, word *string, pos int) bool {
	// Out of board
	if (x < 0 || x >= boardSizeX) || (y < 0 || y >= boardSizeY) {
		return false
	}

	// Get cached data
	ret := &cache[x][y][pos]
	if *ret != none { return *ret == yes }

	// Word not match
	if (*word)[pos] != (*board)[y][x] {
		*ret = no
		return false
	}

	// Found full match
	if pos == len(*word)-1 {
		*ret = yes
		return true
	}

	// Search for every direction
	*ret = no
	for i := 0; i < 8; i++ {
		if hasWord(board, x + dx[i], y + dy[i], word, pos+1) {
			*ret = yes
			return true
		}
	}
	return false
}

func main() {
	var C, N int
	fmt.Scan(&C)
	for ; C > 0; C-- {
		board := make([]string, 5)
		for i := 0; i < boardSizeY; i++ {
			fmt.Scan(&board[i])
		}

		fmt.Scan(&N)
		for ; N > 0; N-- {
			var word string
			fmt.Scan(&word)

			// Initialize cache
			for i := 0; i < boardSizeX; i++ {
				for j := 0; j < boardSizeY; j++ {
					for k := 0; k < wordLengthMax; k++ {
						cache[i][j][k] = none
					}
				}
			}

			// Run
			ok := false
			for y := 0; y < boardSizeY; y++ {
				for x := 0; x < boardSizeX; x++ {
					if hasWord(&board, x, y, &word, 0) {
						ok = true
						break
					}
				}
			}

			if ok {
				fmt.Printf("%s YES\n", word)
			} else {
				fmt.Printf("%s NO\n", word)
			}
		}
	}
}
