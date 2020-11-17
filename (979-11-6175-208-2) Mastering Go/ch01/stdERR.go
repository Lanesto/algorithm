package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = fmt.Sprintf("%s %s", arguments[0], arguments[1])
	}
	io.WriteString(os.Stdout, "This is Standard output\n")
	io.WriteString(os.Stderr, fmt.Sprintf("%s\n", myString))
}