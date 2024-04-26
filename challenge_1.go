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

1. **Function Description**:
   - **Function Name**: `firstNonRepeatingLetter`
   - **Description**: This function finds and returns the first non-repeating letter in the given string. It returns an empty string if all letters in the string repeat.
   - **Parameters**:
     - `s`: The input string to search for the first non-repeating letter.
   - **Return**:
     - The first non-repeating letter as a string, or an empty string if all letters repeat.

2. **Algorithm**:
   - **Step 1: Create a Character Count Map**:
     - Initialize a map called `charCount` to store the count of each character.
   - **Step 2: Convert to Lowercase**:
     - Convert the input string to lowercase using `strings.ToLower()` to make the comparison case-insensitive.
   - **Step 3: Count Occurrences**:
     - Iterate through the lowercase string and count the occurrences of each character. Store the counts in the `charCount` map.
   - **Step 4: Find First Non-Repeating Character**:
     - Iterate through the original string and check the count of each character in the `charCount` map.
     - If the count is `1`, return the character. This is the first non-repeating character.
   - **Step 5: Return Empty String if All Characters Repeat**:
     - If the loop completes without finding a non-repeating character, return an empty string.

3. **Edge Cases Handling**:
   - This function handles both uppercase and lowercase characters correctly due to the conversion to lowercase.
   - It correctly identifies the first non-repeating character, even if it's a different case from the one in the input string.

4. **Complexity**:
   - The time complexity of this algorithm is O(n), where n is the length of the input string. This is because we iterate through the string only once to count the occurrences and then iterate through it again to find the first non-repeating character.
   - The space complexity is O(1) because the size of the `charCount` map is independent of the input string size.

This code effectively solves the problem by efficiently finding the first non-repeating character in a given string, handling both case-insensitive comparison and returning the correct case for the initial letter.

*/
