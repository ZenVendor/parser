package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	parser, err := NewParser(V_LIST, []int{A_OPEN}, map[int]interface{}{})
	if err != nil {
		log.Fatal(err)
	}
	if err := parser.Parse(args); err != nil {
		log.Fatal(err)
	}
	parser.Print()
}
