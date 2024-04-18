package main

import (
	"testing"
	// "fmt
	"go-reloaded/utils"
)

func Test1E(t *testing.T) {
	input := "1E"
	output := utils.HexToDecimal(input)

	if output != "30" {
		t.Fail()
	}
	// fmt.Printf("Input: %s\nOutput: %s\n", input, output)
}
