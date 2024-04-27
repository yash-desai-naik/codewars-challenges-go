package main

func countSubsequences(needle, haystack string) int {
	// Initialize a 2D array to store the counts
	dp := make([][]int, len(needle)+1)
	for i := range dp {
		dp[i] = make([]int, len(haystack)+1)
	}

	// Initialize the first row with 1 (empty needle appears once in any haystack)
	for j := 0; j <= len(haystack); j++ {
		dp[0][j] = 1
	}

	// Fill the dp array using dynamic programming
	for i := 1; i <= len(needle); i++ {
		for j := 1; j <= len(haystack); j++ {
			if needle[i-1] != haystack[j-1] {
				// If characters don't match, copy the count from the previous row
				dp[i][j] = dp[i][j-1]
			} else {
				// If characters match, add counts from both previous row and previous column
				dp[i][j] = (dp[i][j-1] + dp[i-1][j-1]) % 100000000
			}
		}
	}

	// Return the last 8 digits of the count
	return dp[len(needle)][len(haystack)] % 100000000
}

/*
Explanation

**Code Block 1: Initializing a 2D array to store the counts**
```go
dp := make([][]int, len(needle)+1)
for i := range dp {
    dp[i] = make([]int, len(haystack)+1)
}
```
This code creates a 2D array `dp` with dimensions `(len(needle)+1) x (len(haystack)+1)`. The `+1` is because we need to include the empty string (`""` and `""` as well. The `make` function initializes the array with all elements set to 0.

**Code Block 2: Initializing the first row with 1 (empty needle appears once in any haystack)**
```go
for j := 0; j <= len(haystack); j++ {
    dp[0][j] = 1
}
```
This loop sets the first row of the `dp` array to 1. This is because an empty needle (`""` can appear once in any haystack (`haystack`).

**Code Block 3: Filling the dp array using dynamic programming**
```go
for i := 1; i <= len(needle); i++ {
    for j := 1; j <= len(haystack); j++ {
        if needle[i-1] != haystack[j-1] {
            // If characters don't match, copy the count from the previous row
            dp[i][j] = dp[i][j-1]
        } else {
            // If characters match, add counts from both previous row and previous column
            dp[i][j] = (dp[i][j-1] + dp[i-1][j-1]) % 100000000
        }
    }
}
```
This nested loop fills the `dp` array using dynamic programming. The outer loop iterates over each character in the `needle` string, and the inner loop iterates over each character in the `haystack` string.

* If the current characters in the `needle` and `haystack` strings don't match (`needle[i-1] != haystack[j-1]`), the count is copied from the previous row (`dp[i][j-1]`).
* If the characters match, the count is the sum of the counts from the previous row and previous column (`dp[i][j-1] + dp[i-1][j-1]`). The result is taken modulo `100000000` to ensure the count stays within a reasonable range.

**Code Block 4: Returning the last 8 digits of the count**
```go
return dp[len(needle)][len(haystack)] % 100000000
```
The function returns the count stored in the last cell of the `dp` array (`dp[len(needle)][len(haystack)]`) modulo `100000000`. This is done to ensure the count stays within a reasonable range and to avoid integer overflow.


*/
