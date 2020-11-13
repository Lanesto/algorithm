package main

import (
	"fmt"
	"io"
	"os"
	"flag"
)

func readSize(f *os.File, size int) []byte {
	buffer := make([]byte, size)
	n, err := f.Read(buffer)
	if err == io.EOF {
		return nil
	} else if err != nil {
		fmt.Println(err)
		return nil
	}
	return buffer[:n]
}

func main() {
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Usage: -n <int> -f <filename>")
			return
		}
	}()

	n := flag.Int("n", 10, "Number of bytes to read")
	filename := flag.String("f", "", "Name of file to read")
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer f.Close()

	for {
		buf := readSize(f, *n)
		if buf != nil {
			fmt.Print(string(buf))
		} else {
			break
		}
	}
}
