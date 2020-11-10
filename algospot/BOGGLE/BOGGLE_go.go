package main

import "fmt"

// Size of boggle board (5x5)
const BOARD_SIZE_X = 5
const BOARD_SIZE_Y = 5

// Maximum length of word to find given
const WORD_MAX_LENGTH = 10

// Caching for DP
var cache = [BOARD_SIZE_X][BOARD_SIZE_Y][WORD_MAX_LENGTH]byte{}

const (
	NONE  = iota
	FALSE
	TRUE
)

// Possible directions
var dx = [8]int{-1, -1, -1,  0, 0,  1, 1, 1}
var dy = [8]int{-1,  0,  1, -1, 1, -1, 0, 1}

// From position (x, y) can find remaining subword (*word)[pos:]
func hasWord(board *[]string, x, y int, word *string, pos int) bool {
	// Out of board
	if (x < 0 || x >= BOARD_SIZE_X) || (y < 0 || y >= BOARD_SIZE_Y) {
		return false
	}

	// Get cached data
	ret := &cache[x][y][pos]
	if *ret != NONE { return *ret == TRUE }

	// Word not match
	if (*word)[pos] != (*board)[y][x] {
		*ret = FALSE
		return false
	}

	// Found full match
	if pos == len(*word)-1 {
		*ret = TRUE
		return true
	}

	// Search for every direction
	*ret = FALSE
	for i := 0; i < 8; i++ {
		if hasWord(board, x + dx[i], y + dy[i], word, pos+1) {
			*ret = TRUE
			return true
		}
	}
	return false
}

func main() {
	var C, N int
	fmt.Scanf("%d", &C)
	for ; C > 0; C-- {
		board := make([]string, 5)
		for i := 0; i < BOARD_SIZE_Y; i++ { fmt.Scanf("%s", &board[i]) }

		fmt.Scanf("%d", &N)
		for ; N > 0; N-- {
			var word string
			fmt.Scanf("%s", &word)

			// Initialize cache
			for i := 0; i < BOARD_SIZE_X; i++ {
				for j := 0; j < BOARD_SIZE_Y; j++ {
					for k := 0; k < WORD_MAX_LENGTH; k++ {
						cache[i][j][k] = NONE
					}
				}
			}

			// Run
			ok := false
			for y := 0; y < BOARD_SIZE_Y; y++ {
				for x := 0; x < BOARD_SIZE_X; x++ {
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
