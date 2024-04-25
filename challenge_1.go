package main

import (
	"strings"
	"unicode"
)

func firstNonRepeatingLetter(s string) string {
	// Create a map to store the count of each character
	charCount := make(map[rune]int)

	// Convert the string to lowercase for case-insensitive comparison
	lowercaseString := strings.ToLower(s)

	// Count occurrences of each character
	for _, char := range lowercaseString {
		charCount[char]++
	}

	// Find the first non-repeating character
	for _, char := range s {
		if charCount[unicode.ToLower(char)] == 1 {
			return string(char)
		}
	}

	// If all characters are repeating, return an empty string
	return "''"
}

// func run_challenge_1() {
// 	// // Test cases
// 	println(firstNonRepeatingLetter("stress"))   // Output: "t"
// 	println(firstNonRepeatingLetter("sTreSS"))   // Output: "T"
// 	println(firstNonRepeatingLetter("aabbcc"))   // Output: ""
// 	println(firstNonRepeatingLetter("aAbBcCdD")) // Output: ""
// 	println(firstNonRepeatingLetter("aAbBcCdD")) // Output: ""
// 	println(firstNonRepeatingLetter(""))         // Output: ""
// 	println(firstNonRepeatingLetter("abca"))     // Output: "b"
// }
