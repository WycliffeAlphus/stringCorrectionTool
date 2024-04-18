package main

import (
	"fmt"
	"os"
	"go-reloaded/utils2"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Printf("Usage: %s <input file> <output file>\n", os.Args[0])
		os.Exit(1)
	}

	inputFile := args[0]
	outputFile := args[1]

	fileText, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	utils2.ProcessLine(fileText, outputFile)
}
