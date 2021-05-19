package main

import (
	"fmt"
)

type Board [][]int

var (
	boardSize = 5 // 5 * 5

	cellEmpty   = 0 // .
	cellBlocked = 1 // #

	// All possible placement of blocks (4 L-shape, 2 I-shape)
	allBlocks []int

	// 1 for true, 0 for false, -1 for unknown
	cache    = make([]int8, (1 << 25))
	iTrue    = int8(1)
	iFalse   = int8(0)
	iUnknown = int8(-1)
)

// Returns "WINNING" if current turn' player can win else "LOSING"
func blockgame(board Board) string {

	// Returns whether current turn' player can win
	var solve func(int) bool
	solve = func(state int) bool {

		// Retrieve cache
		ret := &cache[state]
		if *ret != iUnknown {
			return *ret == iTrue
		}

		*ret = iFalse
		for _, v := range allBlocks {
			// If block is not placed
			if state&v == 0 {

				// Opponent' defeat -> win
				if !solve(state | v) {
					*ret = iTrue
					return true
				}
			}
		}
		return false
	}

	result := solve(board.toInt())
	if result {
		return "WINNING"
	}
	return "LOSING"
}

func init() {
	// **********************
	// *** Pre-cacluation ***
	// **********************

	// L shaped blocks
	for y := 0; y < (boardSize - 1); y++ {
		for x := 0; x < (boardSize - 1); x++ {
			square := cell(x, y) + cell(x+1, y) + cell(x, y+1) + cell(x+1, y+1)
			allBlocks = append(allBlocks, []int{
				square - cell(x, y),     // ┘
				square - cell(x+1, y),   // └
				square - cell(x, y+1),   // ┐
				square - cell(x+1, y+1), // ┌
			}...)
		}
	}

	// I shaped blocks
	for i := 0; i < (boardSize - 1); i++ {
		for j := 0; j < boardSize; j++ {
			allBlocks = append(allBlocks, []int{
				cell(i, j) + cell(i+1, j), // -
				cell(j, i) + cell(j, i+1), // I
			}...)
		}
	}

	// ******************
	// *** Init cache ***
	// ******************
	for i := 0; i < len(cache); i++ {
		cache[i] = iUnknown
	}

	// Known cases
	cache[0] = iFalse // LOSING
}

func main() {
	var C int
	var row string
	fmt.Scan(&C)
	for ; C > 0; C-- {
		board := make(Board, boardSize)
		for i := 0; i < boardSize; i++ {
			board[i] = make([]int, boardSize)
			fmt.Scan(&row)
			for j, r := range row {
				switch r {
				case '#':
					board[i][j] = cellBlocked
				case '.':
					board[i][j] = cellEmpty
				default:
					panic(fmt.Errorf("unexpected symbol: %v", r))
				}
			}
		}
		result := blockgame(board)
		fmt.Println(result)
	}
}

// Returns identical value for board[y][x]
func cell(x, y int) int {
	return 1 << uint(boardSize*y+x)
}

// Convert board into integer representation
func (b Board) toInt() int {
	acc := 0
	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			if b[y][x] == cellEmpty {
				continue
			}
			acc += cell(x, y)
		}
	}
	return acc
}
