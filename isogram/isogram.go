package isogram

import (
	"strings"
	"unicode"
)

//isogram is a word or phrase without a repeating letter,
//however spaces and hyphens are allowed to appear multiple times.

//IsIsogram checks whether a string is a isogram
func IsIsogram(word string) bool {
	//word is transformed to lowercase
	word = strings.ToLower(word)
	for _, char := range word {
		//checks is a char is a valid letter
		if !unicode.IsLetter(char) {
			continue
		}
		// if char occurance is more than one returns false
		if charCount := strings.Count(word, string(char)); charCount > 1 {
			return false
		}
	}
	return true
}
