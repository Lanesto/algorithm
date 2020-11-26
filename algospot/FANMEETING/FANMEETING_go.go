package main

import (
	"fmt"
	"math/big"
)

// Replace F, M to 0, 1 each
func convert(s string) string {
	ret := make([]byte, len(s))
	for i, c := range s {
		if c == 'M' {
			ret[i] = '1'
		} else {
			ret[i] = '0'
		}
	}
	return string(ret)
}

// Count how many times ALL members and fans hug each other
// For operator X, member X female equals to:
// F X F = Hug        0 & 0 = 0
// F X M = Hug        0 & 1 = 0
// M X F = Hug        1 & 0 = 0
// M X M = Handshake  1 & 1 = 1
func countHugs(members, fans string) int {
	// Convert F, M to 0, 1 for each
	members = convert(members)
	fans = convert(fans)

	// Convert string representation into big.Int
	binMembers, _ := new(big.Int).SetString(members, 2)
	binFans, _ := new(big.Int).SetString(fans, 2)

	// Count hugs
	ret := 0
	zero := big.NewInt(0)

	/* Note: wrong iteration case
	...
	for binFans.Cmp(zero) != 0 {
		...
	}
	...
	// binFans == 0 doesn't mean all fans are gone
	// It could also mean all fans are female
	*/

	// Always len(fans) >= len(members)
	// members 4       [....]    |    [....]
	// fans    7    [.......] -> | -> [.......] (move 3(7-4) times)
	// If from fans does not cover all members then it doesn't necessary to continue
	for i := 0; i < len(fans)-len(members)+1; i++ {
		// members & fans == 0 -> All Hugging
		if new(big.Int).And(binMembers, binFans).Cmp(zero) == 0 {
			ret++
		}
		binFans.Rsh(binFans, 1)
	}
	return ret
}

func main() {
	var C int
	fmt.Scan(&C)
	for ; C > 0; C-- {
		var members, fans string
		fmt.Scan(&members, &fans)

		result := countHugs(members, fans)
		fmt.Println(result)
	}
}
