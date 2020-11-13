package main

import (
	"fmt"
	"bufio"
	"io"
	"os"
	"regexp"
	"net"
)

const IPv4FieldPattern = `(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[0-9]?[0-9])`
const IPv4Pattern = IPv4FieldPattern + "\\." + IPv4FieldPattern + "\\." + IPv4FieldPattern + "\\." + IPv4FieldPattern
var IPv4 = regexp.MustCompile(IPv4Pattern)
var Counter = make(map[string]int)


func CountIPv4All(line string) int {
	match := IPv4.FindAllString(line, -1)
	for _, address := range match {
		trial := net.ParseIP(address)
		if trial.To4() == nil {
			fmt.Printf("Found malformed IPv4 address: %s\n", address)
			continue
		} else {
			Counter[address] += 1
		}
	}
	return len(match)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Requires at least one file to read")
		os.Exit(1)
	}

	for _, filename := range args[1:] {
		f, err := os.Open(args[1])
		if err != nil {
			fmt.Printf("Could not open file %q; skipping\n", filename)
			continue
		}
		defer f.Close()
		reader := bufio.NewReader(f)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("Error occured while reading file %s; error=%s\n", filename, err)
				break
			}

			CountIPv4All(line)
		}
	}

	maxCount := -1
	mostFrequentAddress := ""
	for address, count := range Counter {
		if count > maxCount {
			maxCount = count
			mostFrequentAddress = address
		}
	}
	fmt.Printf("We found most frequent IPv4 address for all given %d files with %d occurences: %s\n", len(args[1:]), maxCount, mostFrequentAddress)
	fmt.Println("* Hint: to test it, use 'grep -o $address $filename | wc -l'")
}
