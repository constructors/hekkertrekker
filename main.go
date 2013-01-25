package main

import (
	"fmt"
	"os"
)

func accept() {
}

func deliver() {
}

func done() {
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing argument")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "accept": accept()
	case "deliver": deliver()
	case "done": done()
	default:
		fmt.Printf("Wrong argument\nUsage: %s [accept|deliver|done]\n", os.Args[0])
		os.Exit(1)
	}
}