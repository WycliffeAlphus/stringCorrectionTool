package utils

import (
	"fmt"
	"strconv"
)
// Function for transforming number from binary to decimal
func BinaryToDecimal(bin string) string {
	decimal, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		panic(err)
	}
	return fmt.Sprint(decimal)
}
