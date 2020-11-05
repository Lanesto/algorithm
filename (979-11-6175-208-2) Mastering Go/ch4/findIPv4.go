package main

import (
	"os"
	"io"
	"bufio"
	"fmt"
	"net"
	"path/filepath"
	"regexp"
)

func findIP(input string) string {
	part := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	pattern := part + "\\." + part + "\\." + part + "\\." + part
	re := regexp.MustCompile(pattern)
	return re.FindString(input)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage: %s logFile\n", filepath.Base(args[0]))
		os.Exit(1)
	}

	for _, filename := range args[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error occured while opening file %s", err)
			os.Exit(-1)
		}
		defer f.Close()
		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Print("Error occured while reading file %s", err)
				break
			}
			ip := findIP(line)
			trial := net.ParseIP(ip)
			if trial.To4() == nil {
				continue
			} else {
				fmt.Println(ip)
			}
		}
	}
}
