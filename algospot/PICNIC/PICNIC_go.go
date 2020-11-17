package main

import "fmt"

const maxStudents = 10

// Values not change during running for single case
var numStudents int
var areFriends [maxStudents][maxStudents]bool
var isTaken [maxStudents]bool

// Make pair who has no partner yet, lowest index first
func countPairings() int {
	// Find lowest index who has no partner yet
	var firstFree int
	for firstFree = 0; firstFree < numStudents; firstFree++ {
		if !isTaken[firstFree] {
			break
		}
	}
	// All students have their partner; found 1 possible combination
	if firstFree == numStudents { return 1 }

	// Find other one who has no pair yet; lowest index first also
	result := 0
	for j := firstFree+1; j < numStudents; j++ {
		if !isTaken[j] && areFriends[firstFree][j] {
			isTaken[firstFree], isTaken[j] = true, true
			result += countPairings()
			isTaken[firstFree], isTaken[j] = false, false
		}
	}
	return result
}

func main() {
	var C int
	fmt.Scan(&C)
	for ; C > 0; C-- {
		var n, m int
		fmt.Scan(&n); fmt.Scan(&m)

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				areFriends[i][j], areFriends[j][i] = false, false
			}
		}
		for k := 0; k < n; k++ {
			isTaken[k] = false
		}

		var x, y int
		for ; m > 0; m-- {
			fmt.Scan(&x); fmt.Scan(&y)
			areFriends[x][y], areFriends[y][x] = true, true
		}

		numStudents = n
		ret := countPairings()
		fmt.Println(ret)
	}
}
