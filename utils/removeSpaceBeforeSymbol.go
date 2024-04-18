package utils

import (
	"regexp"
)

func RemoveSpaceBeforeSymbol(s string) string {
	re := regexp.MustCompile(`(\s+)([,.!:;?])`)
	return re.ReplaceAllString(s, "$2 ") // space after 2 adds a space after replacing the group
}
