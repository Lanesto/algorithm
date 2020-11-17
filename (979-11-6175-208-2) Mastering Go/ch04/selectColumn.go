package main

import (
	"os"
	"bufio"
	"io"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Printf("usage: selectColumn column <file1> [<file2> [....<fileN]]\n")
		os.Exit(1)
	}

	temp, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("column value is not an integer:", temp)
		return
	}

	column := temp
	if column < 0 {
		fmt.Println("Invalid column number!")
		os.Exit(1)
	}

	for _, filename := range arguments[2:] {
		fmt.Println("\t\t", filename)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file %s\n", err)
			continue
		}
		defer f.Close()
		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file %s", err)
			}
			data := strings.Fields(line)
			if len(data) >= column {
				fmt.Println((data[column-1]))
			}
		}
	}
}
