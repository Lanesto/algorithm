package main

import "fmt"

// Board cell type definition
const (
	white = 0
	black = 1
)

// Define relative positions of blocks
// blockShapes[blockType: 0..3][index: 0..2][xy: x=0, y=1]
var blockShapes = [4][3][2]int{
	{ {0, 0},  {1, 0}, {1, 1} },
	{ {0, 0}, {-1, 1}, {0, 1} },
	{ {0, 0},  {0, 1}, {1, 1} },
	{ {0, 0},  {1, 0}, {0, 1} },
}

// Returns whether given position is in range of board
func isInBoard(board *[][]int, x, y int) bool {
	return (y >= 0 && y < len(*board)) && (x >= 0 && x < len((*board)[0]))
}

// Cover selected position with given type of block
// Returns true if covered without overlappings else false
// Argument set is to specify behavior(set: +1, unset: -1)
// to allow duplicated overlapping blocks
func setCover(board *[][]int, x, y, blockType int, set bool) bool {
	// Boolean to check whether block is successfully placed
	// without any of overlapping or out-of-board situations
	blockPlaced := true
	for i := 0; i < 3; i++ {
		nx := x + blockShapes[blockType][i][0]
		ny := y + blockShapes[blockType][i][1]
		if !isInBoard(board, nx, ny) {
			blockPlaced = false
		} else {
			if set {
				(*board)[ny][nx]++
			} else {
				(*board)[ny][nx]--
			}
			// Check whether placed block is overlapping
			if (*board)[ny][nx] > 1 {
				blockPlaced = false
			}
		}
	}
	return blockPlaced
}

// Count number of cases fill given board without any blanks
func countCoverings(board *[][]int) int {
	// Find empty space(white block); i for y, j for x
	x, y := -1, -1
	for i := 0; i < len(*board); i++ {
		for j := 0; j < len((*board)[0]); j++ {
			if (*board)[i][j] == white {
				x, y = j, i
				break
			}
		}
		// Found white block? -> exit nested loop
		if y != -1 {
			break
		}
	}

	// No white block - all filled
	if y == -1 {
		return 1
	}

	// Try each type of blocks
	ret := 0
	for blockType := 0; blockType < 4; blockType++ {
		if setCover(board, x, y, blockType, true) {
			ret += countCoverings(board)
		}
		setCover(board, x, y, blockType, false)
	}
	return ret
}

func main() {
	var C int
	fmt.Scan(&C)
	for ; C > 0; C-- {
		var H, W int
		fmt.Scan(&H, &W)

		board := make([][]int, H)
		for i := 0; i < H; i++ {
			var rawRow string
			fmt.Scan(&rawRow)

			// Convert given string-based board to integer-based one
			row := make([]int, len(rawRow))
			for j, c := range rawRow {
				if c == '.' {
					row[j] = white
				} else {
					row[j] = black
				}
			}
			board[i] = row
		}
		result := countCoverings(&board)
		fmt.Println(result)
	}
}
