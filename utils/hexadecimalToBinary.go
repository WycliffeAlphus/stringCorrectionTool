package utils

import (
	"fmt"
	"strconv"
)

// Function for hexadecimal to decimal
func HexToDecimal(hex string) string {
	decimal, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		panic(err)
	}
	return fmt.Sprint(decimal)
}
