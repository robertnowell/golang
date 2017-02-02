//package word provides utilities for word games
package word

import (
	"unicode"
	// "fmt"
	)
// IsPalindrome reports whether s reads the same way forward and backward
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}

	if len(letters) <= 1 {return false}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	// fmt.Printf("string = '%q'\nletters = %q\n", s, letters)
	return true
}
