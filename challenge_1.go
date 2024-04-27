package main

import (
	"strings"
	"unicode"
)

/**
 * firstNonRepeatingLetter returns the first non-repeating letter in the given string.
 * If all letters repeat, it returns an empty string.
 *
 * @param s The input string to search for the first non-repeating letter.
 * @return The first non-repeating letter as a string, or an empty string if all letters repeat.
 */
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

/*

### Explanation:


**Code Block 1: Creating a map to store the count of each character**
```go
charCount := make(map[rune]int)
```
This line creates a new map (a data structure that stores key-value pairs) called `charCount`. The map will store the count of each character in the input string `s`. The key type of the map is `rune`, which is an alias for `int32` and represents a Unicode code point. The value type of the map is `int`, which represents the count of each character.

**Code Block 2: Converting the string to lowercase for case-insensitive comparison**
```go
lowercaseString := strings.ToLower(s)
```
This line converts the input string `s` to lowercase using the `strings.ToLower` function from the `strings` package. This is done to make the character count comparison case-insensitive.

**Code Block 3: Counting occurrences of each character**
```go
for _, char := range lowercaseString {
    charCount[char]++
}
```
This loop iterates over each character in the lowercase string `lowercaseString`. For each character, it increments the count in the `charCount` map using the syntax `charCount[char]++`. This effectively counts the occurrences of each character in the input string.

**Code Block 4: Finding the first non-repeating character**
```go
for _, char := range s {
    if charCount[unicode.ToLower(char)] == 1 {
        return string(char)
    }
}
```
This loop iterates over each character in the original string `s`. For each character, it checks if the count in the `charCount` map is equal to 1, which means it's a non-repeating character. If it is, the function returns the character as a string using the `string(char)` conversion.

**Code Block 5: Returning an empty string if all characters are repeating**
```go
return "''"
```
If the loop in the previous code block doesn't find a non-repeating character, this line returns an empty string `""`. This is because the input string contains only repeating characters.


*/
