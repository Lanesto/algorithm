package main

import (
	"fmt"
	"os"
	"text/template"
)

type Entry struct {
	Number int
	Square int
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Need only one argument for template file")
		return
	}

	tpl := args[1]
	data := [][]int{{-1, 1}, {-2, 4}, {-3, 9}, {-4, 16}}
	var entries []Entry
	for _, i := range data {
		if len(i) == 2 {
			entries = append(entries, Entry{Number: i[0], Square: i[1]})
		}
	}

	t := template.Must(template.ParseGlob(tpl))
	t.Execute(os.Stdout, entries)
}
