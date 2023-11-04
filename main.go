package main

import (
	"fmt"
	"os"

	"github.com/drc288/codember/src"
)

func main() {
	// Get input from cli
	var listChallenge = []string{"1", "2"}
	argv := os.Args[1:]

	// Check if input is empty

	if len(argv) == 0 {
		fmt.Println("Please provide the codember challenge number, valid options are: ", listChallenge)
		os.Exit(1)
	}

	if argv[0] == "1" {
		src.Challenge1()
	}

	if argv[0] == "2" {
		src.Challenge2()
	}
}