package main

import "fmt"

func main() {
	fmt.Println("*** Challenge 1 ***")
	words := []string{"stress", "sTreSS", "aabbcc", "aAbBcCdD", "aAbBcCdD", "''", "abbca"}
	for _, word := range words {
		fmt.Printf("%-9s: %-10s\n", word, firstNonRepeatingLetter(word))
	}
	fmt.Println()
	fmt.Println("*** Challenge 2 ***")
	needle := "happy birthday"
	haystack := `hhhappyyyy biirrrrrthddaaaayyyyyyy to youuuu
	hhapppyyyy biirtttthdaaay too youuu
	happy birrrthdayy to youuu
	happpyyyy birrtthdaaay tooooo youu`
	fmt.Printf("needle: %s\ncount: %d", needle, countSubsequences(needle, haystack))

	fmt.Println()
	fmt.Println("*** Challenge 3 ***")
	// Creating the example tree
	root := &TreeNode{Value: 17}
	root.Left = &TreeNode{Value: 3}
	root.Left.Left = &TreeNode{Value: 2}
	root.Right = &TreeNode{Value: -10}
	root.Right.Left = &TreeNode{Value: 16}
	root.Right.Right = &TreeNode{Value: 1}
	root.Right.Right.Left = &TreeNode{Value: 13}

	// Calculating the maximum sum of route from root to leaf
	maxSum := maxPathSumFromRootToLeaf(root)
	fmt.Println("Maximum sum of route from root to leaf:", maxSum) // Output: 23
	// Creating the example tree
	root2 := &TreeNode{Value: 5}
	root2.Left = &TreeNode{Value: 4}
	root2.Left.Left = &TreeNode{Value: -80}
	root2.Left.Right = &TreeNode{Value: -60}
	root2.Right = &TreeNode{Value: 10}
	root2.Right.Left = &TreeNode{Value: -90}

	// Calculating the maximum sum of route from root to leaf
	maxSum2 := maxPathSumFromRootToLeaf(root2)
	fmt.Println("Maximum sum of route from root to leaf:", maxSum2) // Output: -51

	fmt.Println()
	fmt.Println("*** Challenge 4 ***")
	// Example 1
	root1 := &TreeNode_{Value: 2}
	root1.Left = &TreeNode_{Value: 8}
	root1.Right = &TreeNode_{Value: 9}
	root1.Left.Left = &TreeNode_{Value: 1}
	root1.Left.Right = &TreeNode_{Value: 3}
	root1.Right.Left = &TreeNode_{Value: 4}
	root1.Right.Right = &TreeNode_{Value: 5}

	fmt.Println(levelOrder(root1)) // Output: [2 8 9 1 3 4 5]

	// Example 2
	root3 := &TreeNode_{Value: 1}
	root3.Left = &TreeNode_{Value: 8}
	root3.Right = &TreeNode_{Value: 4}
	root3.Left.Right = &TreeNode_{Value: 3}
	root3.Right.Right = &TreeNode_{Value: 5}
	root3.Right.Right.Left = &TreeNode_{Value: 7}

	sortedTree := levelOrder(root3)
	fmt.Println(sortedTree) // Output: [1 8 4 3 5 7]
}
