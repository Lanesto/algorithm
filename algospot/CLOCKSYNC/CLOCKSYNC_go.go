package main

import "fmt"

// Number of clocks and switches
const numClocks = 16
const numSwitches = 10

// Infinity
const inf = 1234567890

// switchConn[i] means slice of connected clocks to switch 'i'
var switchConn = [numSwitches][]int{
	{0, 1, 2},
	{3, 7, 9, 11},
	{4, 10, 14, 15},
	{0, 4, 5, 6, 7},
	{6, 7, 8, 10, 12},
	{0, 2, 14, 15},
	{3, 14, 15},
	{4, 5, 7, 14, 15},
	{1, 2, 3, 4, 5},
	{3, 4, 5, 9, 13},
}

// Check whether all clocks are aligned
func isAllClocksAligned(clockState *[]int) bool {
	for _, v := range *clockState {
		if v != 0 {
			return false
		}
	}
	return true
}

// Push switch to change clock state
func pushSwitch(clockState *[]int, whichSwitch int) {
	for _, clock := range switchConn[whichSwitch] {
		temp := ((*clockState)[clock] + 3) % 12
		(*clockState)[clock] = temp
	}
}

// alignClocks returns how many times should we push switches to
// make all clocks to point 12o'clock
// Note: switchToPush must be call with 0 from outside of this function
func alignClocks(clockState *[]int, switchToPush int) int {
	// All switches has been tried
	if switchToPush == numSwitches {
		if isAllClocksAligned(clockState) {
			return 0
		} else {
			return inf
		}
	}

	// Push switches i(0~3) times cuz it will back to original state
	// when pushes 4 times
	ret := inf
	nextSwitch := switchToPush + 1
	for i := 0; i < 4; i++ {
		// Do recursion first (0 times push)
		temp := i + alignClocks(clockState, nextSwitch)
		if temp < ret {
			ret = temp
		}
		// at i == 3, clockState backs to its original state because
		// switch is pushed 4 times
		pushSwitch(clockState, switchToPush)
	}
	return ret
}

func main() {
	var C int
	fmt.Scan(&C)
	for ; C > 0; C-- {
		clockState := make([]int, numClocks)
		for i, _ := range clockState {
			fmt.Scan(&clockState[i])
			// Set 12 to 0 (it is much easier to handle)
			clockState[i] %= 12
		}
		result := alignClocks(&clockState, 0)
		// inf means impossible
		if result >= inf {
			result = -1
		}
		fmt.Println(result)
	}
}
