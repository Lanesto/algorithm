package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"log"
)

type myElement struct {
	Name string
	Surname string
	Id string
}

var DATA = make(map[string]myElement)


func ADD(k string, n myElement) bool {
	if k == "" {
		return false
	}

	if LOOKUP(k) == nil {
		DATA[k] = n
		log.Printf("Added new item with key %s, value %s\n", k, n)
		return true
	}

	return false
}

func DELETE(k string) bool {
	if LOOKUP(k) != nil {
		delete(DATA, k)
		log.Printf("Removed item with key %s\n", k)
		return true
	}
	return false
}

func LOOKUP(k string) *myElement {
	_, ok := DATA[k]
	if ok {
		n := DATA[k]
		log.Printf("Someone looked up key %s\n", k)
		return &n
	} else {
		return nil
	}
}

func CHANGE(k string, n myElement) bool {
	DATA[k] = n
	log.Printf("Item with key %s updated to %s\n", k, n)
	return true
}

func PRINT() {
	for k, v := range DATA {
		fmt.Printf("Key: %s - Value: %s\n", k, v)
	}
}

func main() {
	f, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666) 
	if err != nil {
		fmt.Println("Could not open log file; err=%v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		tokens := strings.Fields(text)
		switch len(tokens) {
		case 0:
			continue
		case 1:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 2:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 3:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 4:
			tokens = append(tokens, "")
		}

		switch tokens[0] {
		case "PRINT":
			PRINT()
		case "STOP":
			return
		case "DELETE":
			if !DELETE(tokens[1]) {
				fmt.Println("Delete operation failed!")
			}
		case "ADD":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !ADD(tokens[1], n) {
				fmt.Println("Add operation failed!")
			}
		case "LOOKUP":
			n := LOOKUP(tokens[1])
			if n != nil {
				fmt.Printf("%v\n", *n)
			}
		case "CHANGE":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !CHANGE(tokens[1], n) {
				fmt.Println("Update operation failed!")
			}
		default:
			fmt.Println("Unknown command - please try again!")
		}
	}
}
