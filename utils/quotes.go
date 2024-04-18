package utils

import (
	"strings"
)

// func main() {
// 	out := checkQuote("I am exactly how they describe me: ' awesome ' me" +
// 	"As Elton John said: ' I am the most well-known homosexual in the world '")
// 	fmt.Printf("\"%s\"", out)
// }

func HandleSingleQuote(s string) string {
	words := strings.Fields(s)
	output := ""
	// fmt.Println(words)
	isStart := true
	for _, w := range words {
		// w == words[i]
		if w == "'" {
			if isStart {
				//fmt.Print(" (start quote) ")
				// fmt.Printf("%s", w)
				output += w
				isStart = false
			} else {
				//fmt.Print(" (end quote) ")
				output += w + " "
				isStart = true
				//fmt.Printf("%s ", w)
			}
		} else {
			// fmt.Printf("%s ", w)
			output += w + " "
		}
	}

	// 
	words = strings.Fields(output)
	output = ""
	for i, w := range words {

		if i == 0 {
			output += w 
		} else if w == "'" {
			output += w
		} else {
			output += " " + w
		}
	}

	return output

}
