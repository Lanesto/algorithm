package main

import "fmt"

// Utility
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Game player
type player int

const (
	player_X player = -1
	player_O        = -player_X
)

// Get opposite player
func (p player) opposite() player {
	return -p
}

// Get cell match this player
func (p player) cell() cell {
	switch p {
	case player_X:
		return cell_X
	case player_O:
		return cell_O
	}
	panic(fmt.Errorf("unexpected player: %v", p))
}

// Convert to readable string
func (p player) toString() string {
	switch p {
	case player_X:
		return "x"
	case player_O:
		return "o"
	}
	panic(fmt.Errorf("unexpected player: %v", p))
}

// Game result
type gameResult int

// Use minInt to get preferred game result among multiple results (win < draw < lose)
const (
	unknown gameResult = iota - 2
	win
	draw
	lose
)

func (g gameResult) invert() gameResult {
	return -g
}

// Game board
type board []row
type row []cell
type cell player

const (
	cell_Empty cell = iota
	cell_X
	cell_O
)

// Confirm all given values are equal each other
func allEqual(cells ...cell) bool {
	if len(cells) == 0 {
		return true
	}

	head := cells[0]
	for _, v := range cells[1:] {
		if head != v {
			return false
		}
		head = v
	}
	return true
}

// Check if player has made any line
func (b board) lineMadeBy(turn player) bool {
	// For rows
	for i := 0; i < 3; i++ {
		if allEqual(append(b[i], turn.cell())...) {
			return true
		}
	}

	// For columns
	for j := 0; j < 3; j++ {
		if allEqual(b[0][j], b[1][j], b[2][j], turn.cell()) {
			return true
		}
	}

	// Two diagonals
	if allEqual(b[0][0], b[1][1], b[2][2], turn.cell()) || allEqual(b[0][2], b[1][1], b[2][0], turn.cell()) {
		return true
	}

	return false
}

// Type alias of function used by forCell()
// Because the cell is passed by pointer, it is possible to make changes to it directly.
type forCellFunc func(x, y int, c *cell) (stop bool)

// Executes given function for each cell of board
// If f returns true, it stops iteration.
func (b board) forCell(f forCellFunc) {
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			if f(x, y, &b[y][x]) {
				return
			}
		}
	}
}

// Convert board to integer for indexing
func (b board) toInt() int {
	ret := 0
	b.forCell(func(x, y int, c *cell) (stop bool) {
		ret *= 3
		ret += int(*c)
		return false
	})
	return ret
}

// Solution
var (
	cache []gameResult
)

// Returns winner of the game TicTacToe. TIE if the game result in draw.
func tictactoe(b board) string {
	// Count all 'x' and 'o' to determine player of current turn.
	cntX, cntO := 0, 0
	b.forCell(func(x, y int, c *cell) (stop bool) {
		switch *c {
		case cell_X:
			cntX++
		case cell_O:
			cntO++
		}
		return false
	})
	turn := player_O
	if cntX <= cntO {
		turn = player_X
	}

	// Recursively calculates player's best game result can be made.
	var canWin func(board, player) gameResult
	canWin = func(b board, p player) gameResult {
		// Retrieve cache
		ret := &cache[b.toInt()]
		if *ret != unknown {
			return *ret
		}

		// There is a line before current player hasn't made any stroke -> Lose
		if b.lineMadeBy(p.opposite()) {
			return lose
		}

		// There is no line but all cells are filled -> Draw
		ay := -1
		b.forCell(func(x, y int, c *cell) (stop bool) {
			if *c == cell_Empty {
				ay = y
				return true // Stop; there's empty cell yet.
			}
			return false // Keep going
		})
		if ay == -1 {
			return draw
		}

		*ret = lose
		b.forCell(func(x, y int, c *cell) (stop bool) {
			if *c == cell_Empty {
				*c = p.cell()
				*ret = gameResult(minInt(
					int(*ret),
					int(canWin(b, p.opposite()).invert())),
				)
				*c = cell_Empty

				// If there are any case player can win
				if *ret == win {
					// Stop iteration (no need to continue)
					return true
				}
			}
			return false // Go on
		})
		return *ret
	}

	// Convert result as problem' spec require
	switch canWin(b, turn) {
	case win:
		return turn.toString()
	case lose:
		return turn.opposite().toString()
	case draw:
		return "TIE"
	}

	panic("unreachable line")
}

func init() {
	// Initialize cache
	cache = make([]gameResult, 19683)
	for i := range cache {
		cache[i] = unknown
	}
}

func main() {
	var C int
	fmt.Scan(&C)
	for ; C > 0; C-- {
		board := make(board, 3)
		for i := 0; i < 3; i++ {
			var s string
			board[i] = make(row, 3)
			fmt.Scan(&s)
			for j, c := range s {
				switch c {
				case 'x':
					board[i][j] = cell_X
				case 'o':
					board[i][j] = cell_O
				default:
					board[i][j] = cell_Empty
				}
			}
		}
		winner := tictactoe(board)
		fmt.Println(winner)
	}
}
