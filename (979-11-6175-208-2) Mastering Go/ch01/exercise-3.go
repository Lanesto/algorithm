package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	f := os.Stdin
	defer f.Close()
	
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "STOP" {
			break
		}

		n, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			fmt.Println("> Non-integer item given!")
		} else {
			fmt.Println(">", n)
		}
	}
	fmt.Println("Stopped")
}
