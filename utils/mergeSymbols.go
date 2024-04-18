package utils

import "regexp"

func MergeSymbols(s string) string {
	re := regexp.MustCompile(`\s+([,.!:;?])`)
	return re.ReplaceAllString(s, "$1")
}