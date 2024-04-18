package utils2

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"go-reloaded/utils"
)

// The function for modifying different lines
func ProcessLine(fileText []byte, outputFile string) {
	// args := os.Args[1:]
	// fileText, _ := os.ReadFile(args[0])
	wordsInText := strings.Split(string(fileText), " ")
	for i := 0; i < (len(wordsInText)); i++ {
		if wordsInText[i] == "(hex)" {
			wordsInText[i-1] = utils.HexToDecimal(wordsInText[i-1])
			wordsInText = utils.RemovePattern(wordsInText, i)
			i--
		}
		if wordsInText[i] == "(bin)" {
			wordsInText[i-1] = utils.BinaryToDecimal(wordsInText[i-1])
			wordsInText = utils.RemovePattern(wordsInText, i)
			i--
		}
		if wordsInText[i] == "(low)" {
			wordsInText[i-1] = strings.ToLower(wordsInText[i-1])
			wordsInText = utils.RemovePattern(wordsInText, i)
			i--
		}
		if wordsInText[i] == "(up)" {
			wordsInText[i-1] = strings.ToUpper(wordsInText[i-1])
			wordsInText = utils.RemovePattern(wordsInText, i)
			i--
		}
		if wordsInText[i] == "(cap)" {
			wordsInText[i-1] = strings.Title(wordsInText[i-1])
			wordsInText = utils.RemovePattern(wordsInText, i)
			i--
		}
		// specified number
		if wordsInText[i] == "(up," {
			digit := strings.Trim(wordsInText[i+1], wordsInText[i+1][1:])
			newdigit, err := strconv.Atoi(digit)
			if err != nil {
				panic(err)
			}

			for j := 1; j <= newdigit; j++ {
				wordsInText[i-j] = strings.ToUpper(wordsInText[i-j])
			}

		}
		if wordsInText[i] == "(low," {
			digit := strings.Trim(wordsInText[i+1], wordsInText[i+1][1:])
			newdigit, err := strconv.Atoi(digit)
			if err != nil {
				panic(err)
			}
			for j := 1; j <= newdigit; j++ {
				wordsInText[i-j] = strings.ToLower(wordsInText[i-j])
			}
		}
		if wordsInText[i] == "(cap," {
			digit := strings.Trim(wordsInText[i+1], wordsInText[i+1][1:])
			newdigit, err := strconv.Atoi(digit)
			if err != nil {
				panic(err)
			}
			for j := 1; j <= newdigit; j++ {
				wordsInText[i-j] = strings.Title(wordsInText[i-j])
			}
		}
	}
	re := regexp.MustCompile(`\(\w+,\s\d+\)`) // regex for command to remove
	byteString := utils.SlicesstringsToBytes(wordsInText)
	output := re.ReplaceAllString(string(byteString), "")     // remove commands after result for the case of specified number
	newOutput := utils.RemoveSpaceBeforeSymbol(output)        // punctuations
	newOutput = utils.MergeSymbols(newOutput)                 // spaces in the middle of punctuations
	newOutput = utils.HandleSingleQuote(newOutput)            // call the created utils package for handling single quotes
	output2 := utils.VowelFunction(strings.Fields(newOutput)) // the function takes an array of strings, thus had to use .fields
	newOutput = strings.Join(output2, " ")                    // change it back to a string
	newOutput2 := newOutput + "\n"
	os.WriteFile(outputFile, []byte(newOutput2), 0o644)
}
