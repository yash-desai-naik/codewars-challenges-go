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

// func run_challenge_2() {
// 	// Test case
// 	needle := "happy birthday"
// 	haystack := `hhhappyyyy biirrrrrthddaaaayyyyyyy to youuuu
// hhapppyyyy biirtttthdaaay too youuu
// happy birrrthdayy to youuu
// happpyyyy birrtthdaaay tooooo youu`

// 	fmt.Println(countSubsequences(needle, haystack)) // Output: 2276800
// }
