package utils

// function for removing the pattern
func RemovePattern(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
