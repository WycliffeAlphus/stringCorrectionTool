package utils

import "strings"

// function for vowels
// first creating boolean  for tracking if the current word has a vowel/h at index zero
func IsVowel(s string) bool {
	vowels := []uint8{'a', 'e', 'i', 'o', 'u', 'h'}
	for _, letter := range vowels {
		if s[0] == letter {
			return true
		}
	}
	return false
}

// now the vowel fuction
// A a
// An an
func VowelFunction(s []string) []string {
	l := len(s)
	for i, word := range s {
		// Next index
		j := i + 1
		if strings.ToLower(word) == "a" && j < l && IsVowel(s[j]) {
			// word can only be A or a
			s[i] = word + "n"
		}
	}
	return s
}
